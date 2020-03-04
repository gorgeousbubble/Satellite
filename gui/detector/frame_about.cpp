/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_about.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "frame_about.h"

// frameAbout UI class

CFrameAboutUI::CFrameAboutUI(CPaintManagerUI* pManager) {
	CDialogBuilder builder;
	CContainerUI* pframeAboutUI = static_cast<CContainerUI*>(builder.Create(_T("frame\\frameAbout.xml"), NULL, NULL, pManager));

	if (pframeAboutUI) {
		this->Add(pframeAboutUI);
	} else {
		this->RemoveAll();
		return;
	}
}
