#include "stdafx.h"
#include "reg.h"

reg::reg()
{
	key = NULL;
	section = HKEY_LOCAL_MACHINE;
}

reg::~reg()
{

}

void reg::check_end_of_string(std::wstring &str)
{
	if (str.size() != 0)
	{
		if (str[str.size() - 1] != 0)
			str.push_back(0);
	}
	else
		str.push_back(0);
}

std::vector<std::wstring> reg::wcs_to_multi_sz(wchar_t *str, unsigned int size)
{
	std::vector<std::wstring> multi_string;
	multi_string.push_back(L"");
	for (unsigned int i = 0, j = 0; i < size; i++)
	{
		multi_string[j] += str[i];
		if (str[i] == 0)
		{
			if (multi_string[j][0] != 0)
			{
				multi_string.push_back(L"");
				j++;
			}
		}
	}
	return multi_string;
}

wchar_t* reg::multi_sz_to_wcs(std::vector<std::wstring> multi_string, unsigned int &size)
{
	size = 0;
	for (unsigned int i = 0; i < multi_string.size(); i++)
		size += multi_string[i].size();
	wchar_t * str = new wchar_t[size];
	for (unsigned int i = 0, j = 0, r = 0; i < size; i++, r++)
	{
		if (r >= multi_string[j].size())
		{
			r = 0;
			j++;
		}
		str[i] = multi_string[j][r];
	}
	return str;
}


void reg::set_section(HKEY t_section)
{
	section = t_section;
}

HKEY reg::get_section()
{
	return section;
}


bool reg::create_key(std::wstring path)
{
	if (RegCreateKeyEx(section, path.c_str(), 0, 0, REG_OPTION_NON_VOLATILE, KEY_ALL_ACCESS, 0, &key, 0) == 0)
	{
		RegCloseKey(key);
		return true;
	}
	return false;
}

bool reg::delete_key(std::wstring path)
{
	if (SHDeleteKey(section, path.c_str())== 0)
		return true;
	return false;
}


bool reg::write_binary_value(std::wstring path, std::wstring name, std::string value)
{
	if (RegCreateKeyEx(section, path.c_str(), 0, 0, REG_OPTION_NON_VOLATILE, KEY_ALL_ACCESS, 0, &key, 0) == 0)
	{
		if (RegSetValueEx(key, name.c_str(), 0, REG_BINARY, (LPBYTE)value.c_str(), value.size()) == 0)
		{
			RegCloseKey(key);
			return true;
		}
		RegCloseKey(key);
	}
	return false;
}

bool reg::write_dword_value(std::wstring path, std::wstring name, long value)
{
	if (RegCreateKeyEx(section, path.c_str(), 0, 0, REG_OPTION_NON_VOLATILE, KEY_ALL_ACCESS, 0, &key, 0) == 0)
	{
		if (RegSetValueEx(key, name.c_str(), 0, REG_QWORD, (LPBYTE)&value, sizeof(value)) == 0)
		{
			RegCloseKey(key);
			return true;
		}
		RegCloseKey(key);
	}
	return false;
}

bool reg::write_qword_value(std::wstring path, std::wstring name, long long value)
{
	if (RegCreateKeyEx(section, path.c_str(), 0, 0, REG_OPTION_NON_VOLATILE, KEY_ALL_ACCESS, 0, &key, 0) == 0)
	{
		if (RegSetValueEx(key, name.c_str(), 0, REG_QWORD, (LPBYTE)&value, sizeof(value)) == 0)
		{
			RegCloseKey(key);
			return true;
		}
		RegCloseKey(key);
	}
	return false;
}

bool reg::write_sz_or_expand_value(std::wstring path, std::wstring name, bool expand, std::wstring value)
{
	DWORD sz_type = REG_SZ;
	if (expand)
		sz_type = REG_EXPAND_SZ;

	this->check_end_of_string(value);

	if (RegCreateKeyEx(section, path.c_str(), 0, 0, REG_OPTION_NON_VOLATILE, KEY_ALL_ACCESS, 0, &key, 0) == 0)
	{
		if (RegSetValueEx(key, name.c_str(), 0, sz_type, (LPBYTE)value.c_str(), value.size() * 2) == 0)
		{
			RegCloseKey(key);
			return true;
		}
		RegCloseKey(key);
	}
	return false;
}

bool reg::write_multi_sz_value(std::wstring path, std::wstring name, std::vector<std::wstring> value)
{
	for (unsigned int i = 0;i < value.size();i++)
		this->check_end_of_string(value[i]);

	if (value.size() != 0)
	{
		if (value[value.size() - 1] != L"")
			value.push_back(L"");
	}
	else
		value.push_back(L"");

	if (RegCreateKeyEx(section, path.c_str(), 0, 0, REG_OPTION_NON_VOLATILE, KEY_ALL_ACCESS, 0, &key, 0) == 0)
	{
		unsigned int size = 0;
		wchar_t *tmp_string = this->multi_sz_to_wcs(value, size);
		if (RegSetValueEx(key, name.c_str(), 0, REG_MULTI_SZ, (LPBYTE)tmp_string, size * 2) == 0)
		{
			delete[] tmp_string;
			RegCloseKey(key);
			return true;
		}
		delete[] tmp_string;
		RegCloseKey(key);
	}
	return false;
}


bool reg::read_binary_value(std::wstring path, std::wstring name, std::string &value)
{
	value.clear();
	if (RegOpenKeyEx(section, path.c_str(), REG_OPTION_OPEN_LINK, KEY_ALL_ACCESS, &key) == 0)
	{
		DWORD size = 0;
		RegQueryValueEx(key, name.c_str(), NULL, NULL, NULL, &size);
		if (size != 0)
		{
			BYTE *tmp_string = new BYTE[size];
			if (RegQueryValueEx(key, name.c_str(), NULL, NULL, (LPBYTE)tmp_string, &size) == 0)
			{
				RegCloseKey(key);
				for (unsigned int i = 0; i < size; i++)
					value.push_back(tmp_string[i]);
				delete[] tmp_string;
				return true;
			}
			delete[] tmp_string;
		}
		RegCloseKey(key);
	}
	return false;
}

bool reg::read_dword_value(std::wstring path, std::wstring name, long &value)
{
	value = 0;
	if (RegOpenKeyEx(section, path.c_str(), REG_OPTION_OPEN_LINK, KEY_ALL_ACCESS, &key) == 0)
	{
		DWORD size = sizeof(value);
		if (RegQueryValueEx(key, name.c_str(), NULL, NULL, (LPBYTE)&value, &size) == 0)
		{
			RegCloseKey(key);
			return true;
		}
		RegCloseKey(key);
	}
	return false;
}

bool reg::read_qword_value(std::wstring path, std::wstring name, long long &value)
{
	value = 0;
	if (RegOpenKeyEx(section, path.c_str(), REG_OPTION_OPEN_LINK, KEY_ALL_ACCESS, &key) == 0)
	{
		DWORD size = sizeof(value);
		if (RegQueryValueEx(key, name.c_str(), NULL, NULL, (LPBYTE)&value, &size) == 0)
		{
			RegCloseKey(key);
			return true;
		}
		RegCloseKey(key);
	}
	return false;
}

bool reg::read_sz_or_expand_value(std::wstring path, std::wstring name, std::wstring &value)
{
	value.clear();
	if (RegOpenKeyEx(section, path.c_str(), REG_OPTION_OPEN_LINK, KEY_ALL_ACCESS, &key) == 0)
	{
		DWORD size = 0;
		RegQueryValueEx(key, name.c_str(), NULL, NULL, NULL, &size);
		if (size % 2 == 0 && size != 0)
		{
			wchar_t *tmp_string = new wchar_t[size / 2];
			if (RegQueryValueEx(key, name.c_str(), NULL, NULL, (LPBYTE)tmp_string, &size) == 0)
			{
				RegCloseKey(key);
				for (unsigned int i = 0;i < size / 2;i++)
					value.push_back(tmp_string[i]);
				delete[] tmp_string;
				return true;
			}
			delete[] tmp_string;
		}
		RegCloseKey(key);
	}
	return false;
}

bool reg::read_multi_sz_value(std::wstring path, std::wstring name, std::vector<std::wstring> &value)
{
	if (RegOpenKeyEx(section, path.c_str(), REG_OPTION_OPEN_LINK, KEY_ALL_ACCESS, &key) == 0)
	{
		DWORD size = 0;
		RegQueryValueEx(key, name.c_str(), NULL, NULL, NULL, &size);
		if (size % 2 == 0 && size != 0)
		{
			wchar_t	 * tmp_string = new wchar_t[size / 2];
			if (RegQueryValueEx(key, name.c_str(), NULL, NULL, (LPBYTE)tmp_string, &size) == 0)
			{
				RegCloseKey(key);
				value = this->wcs_to_multi_sz(tmp_string, size / 2);
				delete[] tmp_string;
				return true;
			}
			delete[] tmp_string;
		}
		RegCloseKey(key);
	}
	return false;
}


bool reg::multi_sz_add_value(std::wstring path, std::wstring name, std::wstring value)
{
	this->check_end_of_string(value);
	std::vector<std::wstring> v;
	long shift = 1;
	if (this->read_multi_sz_value(path, name, v))
	{
		if (v.size() == 1)
			shift = 0;
		std::vector <std::wstring>::iterator it;
		it = v.begin() + shift;
		v.insert(it, value);
		if (this->write_multi_sz_value(path, name, v))
			return true;
	}
	return false;
}

bool reg::multi_sz_delete_value(std::wstring path, std::wstring name, std::wstring value)
{
	this->check_end_of_string(value);
	std::vector<std::wstring> v;
	if (this->read_multi_sz_value(path, name, v))
	{
		unsigned int i = 0;
		bool remove = false;
		for (i; i < v.size(); i++)
		{
			if (v[i] == value)
			{
				remove = true;
				break;
			}
		}
		if (remove)
		{
			v.erase(v.begin() + i);
			std::vector<std::wstring>(v).swap(v);
			if (this->write_multi_sz_value(path, name, v))
				return true;
		}
	}
	return false;
}

bool reg::delete_value(std::wstring path, std::wstring name)
{
	if (RegOpenKeyEx(section, path.c_str(), REG_OPTION_OPEN_LINK, KEY_ALL_ACCESS, &key) == 0)
	{
		if (RegDeleteValue(key, name.c_str()) == 0)
		{
			RegCloseKey(key);
			return true;
		}
		RegCloseKey(key);
	}
	return false;
}




int main()
{
	reg Reg;
	Reg.delete_key(L"SOFTWARE\\TMP");
    return 0;
}

