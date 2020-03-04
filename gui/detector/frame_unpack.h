/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_unpack.h
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#pragma once

#ifndef __FRAME_UNPACK_H_
#define __FRAME_UNPACK_H_

// Include common header files
#include "common_wnd.h"

// Class definition
class CFrameUnpackUI : public CContainerUI {
public:
	CFrameUnpackUI(CPaintManagerUI* pManager);

};

class CFrameUnpackListUI : public IListCallbackUI {
public:
	LPCTSTR GetItemText(CControlUI* pControl, int iIndex, int iSubItem);
};

// External
extern CFrameUnpackListUI g_cFrameUnpackListUI;

#endif // !__FRAME_UNPACK_H_
