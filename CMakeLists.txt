cmake_minimum_required(VERSION 3.16)

include("$ENV{PSPDEV}/psp/share/pspdev.cmake")

project(tinygo_psp_sample C)

option(PSP_KERNEL_MODE "Build the PSP module with kernel attributes" OFF)

set(TARGET_NAME sample)

set(GO_OBJECT "${CMAKE_SOURCE_DIR}/build/goexports.o")
set(PSPSDK_LIBRARIES_FILE "${CMAKE_SOURCE_DIR}/build/pspsdk-libraries.cmake")

if(NOT EXISTS "${PSPSDK_LIBRARIES_FILE}")
    message(FATAL_ERROR
        "Missing ${PSPSDK_LIBRARIES_FILE}; run build-sample.sh so the "
        "PSPSDK libraries can be resolved from goexports.o")
endif()

include("${PSPSDK_LIBRARIES_FILE}")

set(BRIDGE_SOURCES
    ${CMAKE_SOURCE_DIR}/bridge/main.c
    ${CMAKE_SOURCE_DIR}/bridge/library_markers.c
    ${CMAKE_SOURCE_DIR}/bridge/printf.c
)

add_executable(${TARGET_NAME}
    ${BRIDGE_SOURCES}
    ${GO_OBJECT}
)

add_library(pspsdk_go_vlf_printf STATIC
    ${CMAKE_SOURCE_DIR}/bridge/vlf_printf.c
)

set_source_files_properties(${GO_OBJECT} PROPERTIES
    EXTERNAL_OBJECT TRUE
    GENERATED TRUE
)

target_include_directories(${TARGET_NAME} PRIVATE
    ${CMAKE_SOURCE_DIR}
)

if(PSP_KERNEL_MODE)
    target_compile_definitions(${TARGET_NAME} PRIVATE PSPSDK_GO_KERNEL_MODE=1)
endif()

target_link_libraries(${TARGET_NAME}
    -Wl,--start-group
    ${PSPSDK_GO_LIBRARIES}
    pspsdk_go_vlf_printf
    c
    -Wl,--end-group
)

create_pbp_file(
    TARGET ${TARGET_NAME}
    TITLE tinygo-psp-sample
)
