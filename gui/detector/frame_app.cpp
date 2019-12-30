/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_app.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "frame_app.h"

// frameApp UI class

CFrameAppUI::CFrameAppUI(CPaintManagerUI* pManager) {
	CDialogBuilder builder;
	CContainerUI* pframeAppUI = static_cast<CContainerUI*>(builder.Create(_T("frame\\frameApp.xml"), NULL, NULL, pManager));

	if (pframeAppUI) {
		this->Add(pframeAppUI);
	}
	else {
		this->RemoveAll();
		return;
	}
}
