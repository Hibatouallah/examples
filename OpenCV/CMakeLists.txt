# The name of our project is "MYOPENCV". CMakeLists files in this project can 
# refer to the root source directory of the project as ${MYOPENCV_SOURCE_DIR} and 
# to the root binary directory of the project as ${MYOPENCV_BINARY_DIR}. 
cmake_minimum_required(VERSION 2.8)
cmake_policy (VERSION 2.8) 
project (MYOPENCV)

SET(COMPILE_FLAGS "-Wall -Werror -Wno-unused-function")
SET(CMAKE_C_COMPILER "clang")

# Recurse into the subdirectories. This does not actually 
# cause another cmake executable to run. The same process will walk through 
# the project's entire directory structure. 
add_subdirectory (DisplayPicture) 
add_subdirectory (BookExercises) 