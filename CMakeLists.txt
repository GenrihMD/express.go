project (express)
cmake_minimum_required(VERSION 3.2)
include_directories(${CMAKE_JS_INC})
set(
    CMAKE_Go_CREATE_SHARED_LIBRARY
    go build preload -buildmode=c-shared)
file(GLOB SOURCE_FILES "src/*.h" "src/*.go")
add_library(${PROJECT_NAME} SHARED ${SOURCE_FILES})
set_target_properties(${PROJECT_NAME} PROPERTIES PREFIX "" SUFFIX ".node")
target_link_libraries(${PROJECT_NAME} ${CMAKE_JS_LIB})
