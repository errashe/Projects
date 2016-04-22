#include <Windows.h>
#include <string>
#include <vector>
#include <shlwapi.h>
#pragma comment(lib,"Shlwapi.lib")

class reg
{
private:

	HKEY section;

	void check_end_of_string(std::wstring &str);

	std::vector<std::wstring> wcs_to_multi_sz(wchar_t *str, unsigned int size);
	wchar_t* multi_sz_to_wcs(std::vector<std::wstring> multi_string, unsigned int &size);
	
public:

	HKEY key;

	reg();
	~reg();

	void set_section(HKEY t_section);
	HKEY get_section();

	bool create_key(std::wstring path);
	bool delete_key(std::wstring path);

	bool write_binary_value(std::wstring path, std::wstring name, std::string value); // двоичный параметр
	bool write_dword_value(std::wstring path, std::wstring name, long value); // 32 бита (DWORD)
	bool write_qword_value(std::wstring path, std::wstring name, long long value); // 64 бита (QWORD)
	bool write_sz_or_expand_value(std::wstring path, std::wstring name, bool expand, std::wstring value); // строковый или расширяемый строковый
	bool write_multi_sz_value(std::wstring path, std::wstring name, std::vector<std::wstring> value); // мультистроковый параметр

	bool read_binary_value(std::wstring path, std::wstring name, std::string &value);
	bool read_dword_value(std::wstring path, std::wstring name, long &value);
	bool read_qword_value(std::wstring path, std::wstring name, long long &value);
	bool read_sz_or_expand_value(std::wstring path, std::wstring name, std::wstring &value);
	bool read_multi_sz_value(std::wstring path, std::wstring name, std::vector<std::wstring> &value);

	bool multi_sz_add_value(std::wstring path, std::wstring name, std::wstring value);
	bool multi_sz_delete_value(std::wstring path, std::wstring name, std::wstring value);

	bool delete_value(std::wstring path, std::wstring name);
};