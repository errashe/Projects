require 'cmath'
require 'complex'

class DFT
	def self.dftj(vec, j)
		res = Complex(0, 0)
		vec.each_with_index do |e, i|
			res += vec[i] * CMath.exp(Complex(0, (j * i * -2*CMath::PI)/vec.size))
		end
		return res
	end

	def self.dedftj(vec, j)
		res = 0
		vec.each_with_index do |e, i|
			res += vec[i] * CMath.exp(Complex(0, (j * i * 2*CMath::PI)/vec.size))
		end
		return res
	end

	def self.dft(vec)
		return vec.each_index.map { |i| dftj(vec, i) }
	end

	def self.dedft(vec)
		return vec.each_index.map { |i| (dedftj(vec, i)/vec.size).real }
	end
end