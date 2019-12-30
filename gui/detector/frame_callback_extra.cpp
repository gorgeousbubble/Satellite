/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_callback_extra.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "frame_callback_extra.h"

// Dialog Builder Callback Extra Class

CDialogBuilderCallbackEx::CDialogBuilderCallbackEx(CPaintManagerUI* pManager) {
	m_pManager = pManager;
}

CControlUI* CDialogBuilderCallbackEx::CreateControl(LPCTSTR pstrClass)
{
	if (_tcscmp(pstrClass, _T("frameHome")) == 0) {
		return new CFrameHomeUI(m_pManager);
	}

	if (_tcscmp(pstrClass, _T("framePack")) == 0) {
		return new CFramePackUI(m_pManager);
	}

	if (_tcscmp(pstrClass, _T("frameUnpack")) == 0) {
		return new CFrameUnpackUI(m_pManager);
	}

	if (_tcscmp(pstrClass, _T("frameComp")) == 0) {
		return new CFrameCompUI(m_pManager);
	}

	if (_tcscmp(pstrClass, _T("frameDecomp")) == 0) {
		return new CFrameDecompUI(m_pManager);
	}

	return nullptr;
}
