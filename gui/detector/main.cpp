/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		main.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "common.h"

int APIENTRY _tWinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPTSTR lpCmdLine, int nCmdShow) {
	CPaintManagerUI::SetInstance(hInstance);
	CPaintManagerUI::SetResourcePath(CPaintManagerUI::GetInstancePath());

	HANDLE hMutex;
	hMutex = CreateMutex(NULL, TRUE, L"detector");
	if (hMutex)
	{
		if (ERROR_ALREADY_EXISTS == GetLastError())
		{
			return -1;
		}
	}

	HRESULT Hr = ::CoInitialize(NULL);
	if (FAILED(Hr)) return -2;

	CFrameMain cMainFrame;
	cMainFrame.Create(NULL, _T("detector"), UI_WNDSTYLE_FRAME, WS_EX_STATICEDGE | WS_EX_APPWINDOW);
	cMainFrame.CenterWindow();
	cMainFrame.ShowModal();

	::CoUninitialize();

	return 0;
}
