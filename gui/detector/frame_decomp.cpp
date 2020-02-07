/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_decomp.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "frame_decomp.h"

// frameDecomp UI class

CFrameDecompUI::CFrameDecompUI(CPaintManagerUI* pManager) {
	CDialogBuilder builder;
	CContainerUI* pframeDecompUI = static_cast<CContainerUI*>(builder.Create(_T("frame\\frameDecomp.xml"), NULL, NULL, pManager));

	if (pframeDecompUI) {
		this->Add(pframeDecompUI);
	}
	else {
		this->RemoveAll();
		return;
	}
}
