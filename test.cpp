#include <iostream>
#include <string>
#include <array>

#include "helper.cpp"

int main() {
	auto pPipe = ::popen("yes no | cp -i mnt/foo1/* mnt/foo2/ -i 2>&1", "r");
	if (pPipe == nullptr)
		throw std::runtime_error("Cannot open pipe");
	std::array<char, 256> buffer;

	std::string result;
	while (not std::feof(pPipe)) {
		auto bytes = std::fread(buffer.data(), 1, buffer.size(), pPipe);
		result.append(buffer.data(), bytes);
	}

	auto rc = ::pclose(pPipe);
	if (WIFEXITED(rc))
		std::cout << "return status: " << WEXITSTATUS(rc) << std::endl;

	std::cout << result << std::endl;

	return 0;
}
