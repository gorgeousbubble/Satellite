/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_pack.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "frame_pack.h"

// framePack UI class

CFramePackUI::CFramePackUI(CPaintManagerUI* pManager) {
	CDialogBuilder builder;
	CContainerUI* pframePackUI = static_cast<CContainerUI*>(builder.Create(_T("frame\\framePack.xml"), NULL, NULL, pManager));

	if (pframePackUI) {
		this->Add(pframePackUI);
	}
	else {
		this->RemoveAll();
		return;
	}
}
