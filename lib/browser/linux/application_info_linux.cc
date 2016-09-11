#include <string>
#include "base/logging.h"

namespace gallium {

// TODO: move these to brightray and derive them from the running binary if possible
std::string GetApplicationName() { return "Brightray Example"; }
std::string GetApplicationVersion() { return "0.0.0.1"; }

}
