cmake_minimum_required(VERSION 3.16)

include("$ENV{PSPDEV}/psp/share/pspdev.cmake")

project(tinygo_psp_sample C)

set(TARGET_NAME sample)

set(GO_OBJECT "${CMAKE_SOURCE_DIR}/build/goexports.o")

set(BRIDGE_SOURCES
    ${CMAKE_SOURCE_DIR}/bridge/main.c
)

add_executable(${TARGET_NAME}
    ${BRIDGE_SOURCES}
    ${GO_OBJECT}
)

set_source_files_properties(${GO_OBJECT} PROPERTIES
    EXTERNAL_OBJECT TRUE
    GENERATED TRUE
)

target_include_directories(${TARGET_NAME} PRIVATE
    ${CMAKE_SOURCE_DIR}
)

target_link_libraries(${TARGET_NAME}
    pspdebug
    pspdisplay
    pspctrl
    pspge
    c
)

create_pbp_file(
    TARGET ${TARGET_NAME}
    TITLE tinygo-psp-sample
)
