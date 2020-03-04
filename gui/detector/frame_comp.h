/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_comp.h
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#pragma once

#ifndef __FRAME_COMP_H_
#define __FRAME_COMP_H_

// Include common header files
#include "common_wnd.h"

// Class definition
class CFrameCompUI : public CContainerUI {
public:
	CFrameCompUI(CPaintManagerUI* pManager);

};

class CFrameCompListUI : public IListCallbackUI {
public:
	LPCTSTR GetItemText(CControlUI* pControl, int iIndex, int iSubItem);
};

// External
extern CFrameCompListUI g_cFrameCompListUI;

#endif // !__FRAME_COMP_H_
