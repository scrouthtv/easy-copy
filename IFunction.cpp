#include <string>
#include <iostream>

class IFunction {
	private:
		/**
		 * Options are always lowercased
		 */
		virtual void parseOption(std::string option) = 0;

	public:
		void parseOptions(char* argv[]) {

			// ignore 1st and 2nd
			const int argc = sizeof(argv) / sizeof(argv[0]);
			for (int i = 2; i < argc; i++) {

				std::cout << argv[i] << std::endl;
				// if is an option, remove -(-) and lowercase it, then try catch the unknown option etc etc
			}
		}
		virtual int run() = 0;
};

class CPFunction: public IFunction {
	private:
		bool archive = false;
		bool force = false;

	public:
		void parseOption(std::string option) {
			if (option.compare("a") == 0 || option.compare("archive") == 0) archive = true;
			else if (option.compare("f") == 0 || option.compare("force") == 0) force = true;
			else throw UnknownOptionException();
		}

		int run() {
			std::cout << "bla";
			return 0;
		}
};

class MVFunction : public IFunction {
	public:
		void parseOption(std::string option) {

		}

		int run() {
			return 0;
		}
};

class RMFunction : public IFunction {
	public:
		void parseOption(std::string option) {

		}

		int run() {
			return 0;
		}
};
