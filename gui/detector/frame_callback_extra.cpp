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
#include "frame_pack.h"
#include "frame_unpack.h"
#include "frame_comp.h"
#include "frame_decomp.h"
#include "frame_base.h"
#include "frame_app.h"
#include "frame_more.h"
#include "frame_about.h"

// Dialog Builder Callback Extra Class

CDialogBuilderCallbackEx::CDialogBuilderCallbackEx(CPaintManagerUI* pManager) {
	m_pManager = pManager;
}

CControlUI* CDialogBuilderCallbackEx::CreateControl(LPCTSTR pstrClass)
{
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

	if (_tcscmp(pstrClass, _T("frameBase")) == 0) {
		return new CFrameBaseUI(m_pManager);
	}

	if (_tcscmp(pstrClass, _T("frameApp")) == 0) {
		return new CFrameAppUI(m_pManager);
	}

	if (_tcscmp(pstrClass, _T("frameMore")) == 0) {
		return new CFrameMoreUI(m_pManager);
	}

	if (_tcscmp(pstrClass, _T("frameAbout")) == 0) {
		return new CFrameAboutUI(m_pManager);
	}

	return nullptr;
}
