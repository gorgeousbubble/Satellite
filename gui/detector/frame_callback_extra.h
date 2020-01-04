/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_callback_extra.h
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#pragma once

#ifndef __FRAME_CALLBACK_EXTRA_H_
#define __FRAME_CALLBACK_EXTRA_H_

// Include DUI frame header files
#include "common_wnd.h"

// Class definition
class CDialogBuilderCallbackEx : public IDialogBuilderCallback {
protected:
	CPaintManagerUI* m_pManager;

public:
	CDialogBuilderCallbackEx(CPaintManagerUI* pManager);
	CControlUI* CreateControl(LPCTSTR pstrClass);

};

#endif // !__FRAME_CALLBACK_EXTRA_H_
