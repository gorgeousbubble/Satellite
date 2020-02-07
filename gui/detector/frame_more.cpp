/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_more.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "frame_more.h"

// frameMore UI class

CFrameMoreUI::CFrameMoreUI(CPaintManagerUI* pManager) {
	CDialogBuilder builder;
	CContainerUI* pframeMoreUI = static_cast<CContainerUI*>(builder.Create(_T("frame\\frameMore.xml"), NULL, NULL, pManager));

	if (pframeMoreUI) {
		this->Add(pframeMoreUI);
	}
	else {
		this->RemoveAll();
		return;
	}
}
