/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		types.h
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#pragma once

#ifndef __TYPES_H_
#define __TYPES_H_

#include <Windows.h>

// type pack
typedef struct {
	int nSerial;
	char chName[MAX_PATH];
	char chPath[MAX_PATH];
} TPacketInfo, *LPTPacketInfo;

#endif // !__TYPES_H_
