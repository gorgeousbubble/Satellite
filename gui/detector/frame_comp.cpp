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
#include "frame_main.h"

// frameComp UI class
CFrameCompListUI g_cFrameCompListUI;

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

LPCTSTR CFrameCompListUI::GetItemText(CControlUI* pControl, int iIndex, int iSubItem) {
	TCHAR szBuf[MAX_PATH] = { 0 };

	switch (iSubItem) {
	case 0: {
		_stprintf_s(szBuf, _T("%d"), g_pFrameMain->m_vecComp[iIndex].nSerial);
	}
		  break;
	case 1: {
		int iLen = sizeof(g_pFrameMain->m_vecComp[iIndex].chName);
		LPWSTR lpText = new WCHAR[iLen + 1];
		::ZeroMemory(lpText, (iLen + 1) * sizeof(WCHAR));
		::MultiByteToWideChar(CP_ACP, 0, g_pFrameMain->m_vecComp[iIndex].chName, -1, (LPWSTR)lpText, iLen);
		_stprintf_s(szBuf, lpText);
		delete[] lpText;
	}
		  break;
	case 2: {
		int iLen = sizeof(g_pFrameMain->m_vecComp[iIndex].chPath);
		LPWSTR lpText = new WCHAR[iLen + 1];
		::ZeroMemory(lpText, (iLen + 1) * sizeof(WCHAR));
		::MultiByteToWideChar(CP_ACP, 0, g_pFrameMain->m_vecComp[iIndex].chPath, -1, (LPWSTR)lpText, iLen);
		_stprintf_s(szBuf, lpText);
		delete[] lpText;
	}
		  break;
	}

	pControl->SetUserData(szBuf);
	return pControl->GetUserData();
}
