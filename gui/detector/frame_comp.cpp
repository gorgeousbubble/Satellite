/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_comp.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "frame_comp.h"

// frameComp UI class

CFrameCompUI::CFrameCompUI(CPaintManagerUI* pManager) {
	CDialogBuilder builder;
	CContainerUI* pframeCompUI = static_cast<CContainerUI*>(builder.Create(_T("frame\\frameComp.xml"), NULL, NULL, pManager));

	if (pframeCompUI) {
		this->Add(pframeCompUI);
	}
	else {
		this->RemoveAll();
		return;
	}
}
