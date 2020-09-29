#include <iostream>
#include <string>
#include <exception>

#include "helper.cpp"
#include "IFunction.cpp"

enum UsageMessage { mainUsage, cpUsage, mvUsage, rmUsage };

int err_usage(UsageMessage message, char* argv[]);

int main(int argc, char* argv[]) {

	IFunction* func;

	if (argc < 2) return err_usage(mainUsage, argv);
	else if (iequals(argv[1], "cp")) func = new CPFunction;
	else if (iequals(argv[1], "mv")) func = new CPFunction;
	else if (iequals(argv[1], "rm")) func = new RMFunction;
	else return err_usage(mainUsage, argv);

	try {
		func -> parseOptions(argv);
	} catch (UnknownOptionException &ex) {
		std::cout << ex.what() << std::endl;
	}
	func -> run();
}

int err_usage(UsageMessage message, char* argv[]) {
	switch (message) {
		case mainUsage:
			std::cout << "Usage: " << std::endl;
			err_usage(cpUsage, argv);
			err_usage(mvUsage, argv);
			err_usage(rmUsage, argv);
			break;

		case cpUsage: std::cout << argv[0] << " cp SOURCE ... DEST" << std::endl; break;
		case mvUsage: std::cout << argv[0] << " mv SOURCE ... DEST" << std::endl; break;
		case rmUsage: std::cout << argv[0] << " rm FILE ..." << std::endl; break;
	}

	return 255;
}

