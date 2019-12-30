/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_base.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "frame_base.h"

// frameBase UI class

CFrameBaseUI::CFrameBaseUI(CPaintManagerUI* pManager) {
	CDialogBuilder builder;
	CContainerUI* pframeBaseUI = static_cast<CContainerUI*>(builder.Create(_T("frame\\frameBase.xml"), NULL, NULL, pManager));

	if (pframeBaseUI) {
		this->Add(pframeBaseUI);
	}
	else {
		this->RemoveAll();
		return;
	}
}
