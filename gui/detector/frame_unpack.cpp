/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_unpack.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "frame_unpack.h"

// frameUnpack UI class

CFrameUnpackUI::CFrameUnpackUI(CPaintManagerUI* pManager) {
	CDialogBuilder builder;
	CContainerUI* pframeUnpackUI = static_cast<CContainerUI*>(builder.Create(_T("frame\\frameUnpack.xml"), NULL, NULL, pManager));

	if (pframeUnpackUI) {
		this->Add(pframeUnpackUI);
	}
	else {
		this->RemoveAll();
		return;
	}
}
