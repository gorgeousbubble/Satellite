/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		common.h
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#pragma once

#ifndef __COMMON_H_
#define __COMMON_H_

// Include Windows header files
#include <Windows.h>
#include <Shlwapi.h>
#include <CommCtrl.h>

// Include C standard header files
#include <stdio.h>
#include <stdlib.h>
#include <malloc.h>
#include <memory.h>
#include <mmreg.h>
#include <wchar.h>
#include <tchar.h>
#include <time.h>
#include <mmsystem.h>

// Include C++ standard header files
#include <iostream>
#include <vector>

// Include DUI header files
#include "common_wnd.h"
#include "frame_callback_extra.h"
#include "frame_main.h"
#include "frame_pack.h"
#include "frame_unpack.h"
#include "frame_comp.h"
#include "frame_decomp.h"
#include "frame_base.h"
#include "frame_app.h"
#include "frame_more.h"
#include "frame_about.h"

// Include resource header files
#include "resource.h"

// Include Statics libraries
#pragma comment(lib, "Shlwapi.lib")

#endif // !__COMMON_H_
