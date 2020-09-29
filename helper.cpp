#include <exception>
#include <string>

struct UnknownOptionException : public std::exception {
	const char * what () const throw () {
		return "Unknown option: ...";
	}
};

bool iequals(const std::string& a, const std::string& b) {
	return std::equal(a.begin(), a.end(), b.begin(), b.end(),
			[](char a, char b) {
			return tolower(a) == tolower(b);
			});
	// see https://stackoverflow.com/questions/11635/case-insensitive-string-comparison-in-c
}
