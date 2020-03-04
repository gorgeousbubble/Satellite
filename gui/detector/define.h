/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		define.h
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#pragma once

#ifndef __DEFINE_H_
#define __DEFINE_H_

// Disable warning
#pragma warning (disable:4996)

// Timer
#define TIM_PROGRESS_REFRESH_PACKET				0
#define TIM_PROGRESS_REFRESH_UNPACK				1
#define TIM_PROGRESS_REFRESH_COMP               2
#define TIM_PROGRESS_REFRESH_DECOMP             3

// Constants
#define CONST_PROGRESS_REFRESH_TIME				100

// Messages
#define WM_USER_MESSAGE_MENU              (WM_USER + 1)

#define WM_USER_MESSAGE_PACKET_SEARCH     (WM_USER + 10)
#define WM_USER_MESSAGE_PACKET_ADDITEM    (WM_USER + 11)
#define WM_USER_MESSAGE_PACKET_RESULT     (WM_USER + 12)
#define WM_USER_MESSAGE_UNPACK_SEARCH     (WM_USER + 13)
#define WM_USER_MESSAGE_UNPACK_ADDITEM    (WM_USER + 14)
#define WM_USER_MESSAGE_UNPACK_RESULT     (WM_USER + 15)
#define WM_USER_MESSAGE_COMP_SEARCH       (WM_USER + 16)
#define WM_USER_MESSAGE_COMP_ADDITEM      (WM_USER + 17)
#define WM_USER_MESSAGE_COMP_RESULT       (WM_USER + 18)
#define WM_USER_MESSAGE_DECOMP_SEARCH     (WM_USER + 19)
#define WM_USER_MESSAGE_DECOMP_ADDITEM    (WM_USER + 20)
#define WM_USER_MESSAGE_DECOMP_RESULT     (WM_USER + 21)

#endif // !__DEFINE_H_
