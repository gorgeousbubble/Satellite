/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_pack.h
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#pragma once

#ifndef __FRAME_PACK_H_
#define __FRAME_PACK_H_

// Include common header files
#include "common_wnd.h"

// Class definition
class CFramePackUI : public CContainerUI {
public:
	CFramePackUI(CPaintManagerUI* pManager);

};

class CFramePackListUI : public IListCallbackUI {
public:
	LPCTSTR GetItemText(CControlUI* pControl, int iIndex, int iSubItem);
};

// External
extern CFramePackListUI g_cFramePackListUI;

#endif // !__FRAME_PACK_H_
