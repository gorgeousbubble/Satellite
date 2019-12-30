/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_home.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "frame_home.h"

// frameHome UI class

CFrameHomeUI::CFrameHomeUI(CPaintManagerUI* pManager) {
	CDialogBuilder builder;
	CContainerUI* pframeHomeUI = static_cast<CContainerUI*>(builder.Create(_T("frame\\frameHome.xml"), NULL, NULL, pManager));

	if (pframeHomeUI) {
		this->Add(pframeHomeUI);
	} else {
		this->RemoveAll();
		return;
	}
}
