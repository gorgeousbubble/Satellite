/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_main.cpp
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#include "common.h"
#include "define.h"

CFrameMain* g_pFrameMain;

//----------------------------------------------
// @Function:	CurlRequestReply()
// @Purpose: CFrameMain curl request reply
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
size_t CurlRequestReply(void* ptr, size_t size, size_t nmemb, void* stream) {
	cout << "----->reply" << endl;
	string* str = (string*)stream;
	cout << *str << endl;
	(*str).append((char*)ptr, size * nmemb);
	return size * nmemb;
}

//----------------------------------------------
// @Function:	CurlGetRequest()
// @Purpose: CFrameMain curl get request
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CURLcode CurlGetRequest(const string& url, string& response) {
	CURL* curl = curl_easy_init();
	CURLcode res;

	if (curl)
	{
		curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
		curl_easy_setopt(curl, CURLOPT_SSL_VERIFYPEER, false);
		curl_easy_setopt(curl, CURLOPT_SSL_VERIFYHOST, false);
		curl_easy_setopt(curl, CURLOPT_VERBOSE, 1);
		curl_easy_setopt(curl, CURLOPT_READFUNCTION, NULL);
		curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, CurlRequestReply);
		curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void*)&response);
		curl_easy_setopt(curl, CURLOPT_NOSIGNAL, 1);
		curl_easy_setopt(curl, CURLOPT_HEADER, 1);
		curl_easy_setopt(curl, CURLOPT_CONNECTTIMEOUT, 3);
		curl_easy_setopt(curl, CURLOPT_TIMEOUT, 120);

		res = curl_easy_perform(curl);
	}

	curl_easy_cleanup(curl);
	return res;
}

//----------------------------------------------
// @Function:	CurlPostRequest()
// @Purpose: CFrameMain curl post request
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CURLcode CurlPostRequest(const string& url, const string& data, string& response) {
	CURL* curl = curl_easy_init();
	CURLcode res;

	if (curl)
	{
		curl_easy_setopt(curl, CURLOPT_POST, 1);
		curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
		curl_easy_setopt(curl, CURLOPT_POSTFIELDS, data.c_str());
		curl_easy_setopt(curl, CURLOPT_SSL_VERIFYPEER, false);
		curl_easy_setopt(curl, CURLOPT_SSL_VERIFYHOST, false);
		curl_easy_setopt(curl, CURLOPT_VERBOSE, 1);
		curl_easy_setopt(curl, CURLOPT_READFUNCTION, NULL);
		curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, CurlRequestReply);
		curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void*)&response);
		curl_easy_setopt(curl, CURLOPT_NOSIGNAL, 1);
		curl_easy_setopt(curl, CURLOPT_HEADER, 1);
		curl_easy_setopt(curl, CURLOPT_CONNECTTIMEOUT, 3);
		curl_easy_setopt(curl, CURLOPT_TIMEOUT, 120);

		res = curl_easy_perform(curl);
	}

	curl_easy_cleanup(curl);
	return res;
}

//----------------------------------------------
// @Function:	GetWindowClassName()
// @Purpose: CFrameMain get window class name
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LPCTSTR CFrameMain::GetWindowClassName() const {
	return _T("detector");
}

//----------------------------------------------
// @Function:	GetClassStyle()
// @Purpose: CFrameMain get class style
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
UINT CFrameMain::GetClassStyle() const {
	return CS_DBLCLKS;
}

//----------------------------------------------
// @Function:	Notify()
// @Purpose: CFrameMain window messages notify
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::Notify(TNotifyUI& msg) {
	if (msg.sType == _T("click")) {
		if (msg.pSender == m_pCloseBtn) {
			OnLButtonClickedCloseBtn();
		} else if (msg.pSender == m_pRestoreBtn) {
			OnLButtonClickedRestoreBtn();
		} else if (msg.pSender == m_pMaxBtn) {
			OnLButtonClickedMaxBtn();
		} else if (msg.pSender == m_pMinBtn) {
			OnLButtonClickedMinBtn();
		} else if (msg.pSender == m_pPackAddBtn) {
			OnLButtonClickedPacketAddBtn();
		} else if (msg.pSender == m_pPackDelBtn) {
			OnLButtonClickedPacketDelBtn();
		} else if (msg.pSender == m_pPackClrBtn) {
			OnLButtonClickedPacketClrBtn();
		} else if (msg.pSender == m_pPackExportBtn) {
			OnLButtonClickedPacketExportBtn();
		} else if (msg.pSender == m_pPackStartBtn) {
			OnLButtonClickedPacketStartBtn();
		} else if (msg.pSender == m_pUnpackSelBtn) {
			OnLButtonClickedUnpackSelectBtn();
		} else if (msg.pSender == m_pUnpackUnselBtn) {
			OnLButtonClickedUnpackUnselectBtn();
		} else if (msg.pSender == m_pUnpackRetBtn) {
			OnLButtonClickedUnpackRestrictBtn();
		} else if (msg.pSender == m_pUnpackUnretBtn) {
			OnLButtonClickedUnpackUnrestrictBtn();
		} else if (msg.pSender == m_pUnpackDetBtn) {
			OnLButtonClickedUnpackDetialBtn();
		} else if (msg.pSender == m_pUnpackImportBtn) {
			OnLButtonClickedUnpackImportBtn();
		} else if (msg.pSender == m_pUnpackExportBtn) {
			OnLButtonClickedUnpackExportBtn();
		} else if (msg.pSender == m_pUnpackStartBtn) {
			OnLButtonClickedUnpackStartBtn();
		} else if (msg.pSender == m_pCompAddBtn) {
			OnLButtonClickedCompAddBtn();
		} else if (msg.pSender == m_pCompDelBtn) {
			OnLButtonClickedCompDelBtn();
		} else if (msg.pSender == m_pCompClrBtn) {
			OnLButtonClickedCompClrBtn();
		} else if (msg.pSender == m_pCompExportBtn) {
			OnLButtonClickedCompExportBtn();
		} else if (msg.pSender == m_pCompStartBtn) {
			OnLButtonClickedCompStartBtn();
		} else if (msg.pSender == m_pDecompDetBtn) {
			OnLButtonClickedDecompDetialBtn();
		} else if (msg.pSender == m_pDecompImportBtn) {
			OnLButtonClickedDecompImportBtn();
		} else if (msg.pSender == m_pDecompExportBtn) {
			OnLButtonClickedDecompExportBtn();
		} else if (msg.pSender == m_pDecompStartBtn) {
			OnLButtonClickedDecompStartBtn();
		}
	} else if (msg.sType == _T("selectchanged")) {
		if (msg.pSender == m_pPackOpt) {
			m_pMainTab->SelectItem(0);
		} else if (msg.pSender == m_pUnpackOpt) {
			m_pMainTab->SelectItem(1);
		} else if (msg.pSender == m_pCompOpt) {
			m_pMainTab->SelectItem(2);
		} else if (msg.pSender == m_pDecompOpt) {
			m_pMainTab->SelectItem(3);
		} else if (msg.pSender == m_pBaseOpt) {
			m_pMainTab->SelectItem(4);
		} else if (msg.pSender == m_pAppOpt) {
			m_pMainTab->SelectItem(5);
		} else if (msg.pSender == m_pMoreOpt) {
			m_pMainTab->SelectItem(6);
		} else if (msg.pSender == m_pAboutOpt) {
			m_pMainTab->SelectItem(7);
		}
	} else if (msg.sType == _T("textchanged")) {

	}
}

//----------------------------------------------
// @Function:	HandleMessage()
// @Purpose: CFrameMain handle window messages
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::HandleMessage(UINT uMsg, WPARAM wParam, LPARAM lParam) {
	LRESULT lRes = 0;
	BOOL bHandled = TRUE;

	switch (uMsg) {
	case WM_CREATE:
		lRes = OnCreate(uMsg, wParam, lParam, bHandled);
		break;
	case WM_LBUTTONDOWN:
		lRes = OnLButtonDown(uMsg, wParam, lParam, bHandled);
		break;
	case WM_CLOSE:
		lRes = OnClose(uMsg, wParam, lParam, bHandled);
		break;
	case WM_TIMER:
		lRes = OnTimer(uMsg, wParam, lParam, bHandled);
		break;
	case WM_NCACTIVATE:
		lRes = OnNcActivate(uMsg, wParam, lParam, bHandled);
		break;
	case WM_NCCALCSIZE:
		lRes = OnNcCalcSize(uMsg, wParam, lParam, bHandled);
		break;
	case WM_NCPAINT:
		lRes = OnNcPaint(uMsg, wParam, lParam, bHandled);
		break;
	case WM_NCHITTEST:
		lRes = OnNcHitTest(uMsg, wParam, lParam, bHandled);
		break;
	case WM_SIZE:
		lRes = OnSize(uMsg, wParam, lParam, bHandled);
		break;
	case WM_GETMINMAXINFO:
		lRes = OnGetMinMaxInfo(uMsg, wParam, lParam, bHandled);
		break;
	case WM_SYSCOMMAND:
		lRes = OnSysCommand(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_MENU:
		lRes = OnUserMessageMenu(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_PACKET_SEARCH:
		lRes = OnUserMessagePacketSearch(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_PACKET_ADDITEM:
		lRes = OnUserMessagePacketAddItem(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_PACKET_RESULT:
		lRes = OnUserMessagePacketResult(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_UNPACK_SEARCH:
		lRes = OnUserMessageUnpackSearch(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_UNPACK_ADDITEM:
		lRes = OnUserMessageUnpackAddItem(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_UNPACK_RESULT:
		lRes = OnUserMessageUnpackResult(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_COMP_SEARCH:
		lRes = OnUserMessageCompSearch(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_COMP_ADDITEM:
		lRes = OnUserMessageCompAddItem(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_COMP_RESULT:
		lRes = OnUserMessageCompResult(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_DECOMP_SEARCH:
		lRes = OnUserMessageDecompSearch(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_DECOMP_ADDITEM:
		lRes = OnUserMessageDecompAddItem(uMsg, wParam, lParam, bHandled);
		break;
	case WM_USER_MESSAGE_DECOMP_RESULT:
		lRes = OnUserMessageDecompResult(uMsg, wParam, lParam, bHandled);
		break;
	default:
		bHandled = FALSE;
		break;
	}

	if (bHandled) {
		return lRes;
	}

	if (m_PaintManager.MessageHandler(uMsg, wParam, lParam, lRes)) {
		return lRes;
	}

	return CWindowWnd::HandleMessage(uMsg, wParam, lParam);
}

//----------------------------------------------
// @Function:	OnCreate()
// @Purpose: CFrameMain create window message
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnCreate(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	m_PaintManager.Init(m_hWnd);

	CDialogBuilder builder;
	CDialogBuilderCallbackEx cb(&GetPaintManager());
	CControlUI* pRoot = builder.Create(_T("frame\\frameMain.xml"), (LPCTSTR)0, &cb, &m_PaintManager);
	ASSERT(pRoot && "Failed to parse XML");

	m_PaintManager.AttachDialog(pRoot);
	m_PaintManager.AddNotifier(this);

	ConstructExtra();
	InitMenuShow();
	InitWindowSharp();
	InitControls();
	InitCurl();

	return 0;
}

//----------------------------------------------
// @Function:	OnLButtonDown()
// @Purpose: CFrameMain mouse click left button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnLButtonDown(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	bHandled = FALSE;
	return 0;
}

//----------------------------------------------
// @Function:	OnClose()
// @Purpose: CFrameMain close window message
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnClose(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	Shell_NotifyIcon(NIM_DELETE, &m_nid);

	DestructExtra();

	bHandled = FALSE;
	return 0;
}

//----------------------------------------------
// @Function:	OnTimer()
// @Purpose: CFrameMain timer trigger
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnTimer(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	DWORD nID = wParam;

	USES_CONVERSION;
	
	// switch timer id...
	switch (nID) {
	case TIM_PROGRESS_REFRESH_PACKET: {
		CDuiString strPacketType = m_pPackTypeEdt->GetText();
		CDuiString strPacketJson;
		strPacketJson = SplicePackProcessRequestJson(strPacketType);

		// send post pack request...
		string url;
		string result;
		CURLcode resp;

		url = "http://localhost:8080/satellite/pack/p";
		resp = CurlPostRequest(url, string(T2A(strPacketJson.GetData())), result);
		if (resp != CURLE_OK) {
			::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_PACKET);
			MessageBoxA(this->GetHWND(), "请求打包进度失败!", "警告", MB_OK | MB_ICONWARNING);
		}

		if (result.empty()) {
			::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_PACKET);
			MessageBoxA(this->GetHWND(), "请求打包进度失败!", "警告", MB_OK | MB_ICONWARNING);
			return 0;
		}

		// find number done & work...
		string done;
		string work;
		int res;

		// get done value...
		res = GetValueFromResponse(result, "\"done\": ", ",", done);
		if (res != 0) {
			::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_PACKET);
			MessageBoxA(this->GetHWND(), "获取打包进度失败!", "警告", MB_OK | MB_ICONWARNING);
			return 0;
		}
		// get work value...
		res = GetValueFromResponse(result, "\"work\": ", "\n", work);
		if (res != 0) {
			::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_PACKET);
			MessageBoxA(this->GetHWND(), "获取打包进度失败!", "警告", MB_OK | MB_ICONWARNING);
			return 0;
		}

		DWORD dwDone = StringToDword(done);
		DWORD dwWork = StringToDword(work) / 128;

		// update process information...
		int value = (int)((float)dwDone * 100 / (float)dwWork);
		if (value >= 100) {
			value = 100;
			::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_PACKET);
		}
		m_pPackProgress->SetValue(value);

		break;
	}
	case TIM_PROGRESS_REFRESH_UNPACK: {
		CDuiString strUnpackJson;
		CDuiString strUnpackSrc;
		
		strUnpackSrc = m_pUnpackSrcEdt->GetText();
		strUnpackJson = SpliceUnpackProcessRequestJson(strUnpackSrc);

		// send post pack request...
		string url;
		string result;
		CURLcode resp;

		url = "http://localhost:8080/satellite/unpack/p";
		resp = CurlPostRequest(url, string(T2A(strUnpackJson.GetData())), result);
		if (resp != CURLE_OK) {
			::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_UNPACK);
			MessageBoxA(this->GetHWND(), "请求拆包进度失败!", "警告", MB_OK | MB_ICONWARNING);
		}

		if (result.empty()) {
			::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_UNPACK);
			MessageBoxA(this->GetHWND(), "请求拆包进度失败!", "警告", MB_OK | MB_ICONWARNING);
			return 0;
		}

		// find number done & work...
		string done;
		string work;
		int res;

		// get done value...
		res = GetValueFromResponse(result, "\"done\": ", ",", done);
		if (res != 0) {
			::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_UNPACK);
			MessageBoxA(this->GetHWND(), "请求拆包进度失败!", "警告", MB_OK | MB_ICONWARNING);
			return 0;
		}
		// get work value...
		res = GetValueFromResponse(result, "\"work\": ", "\n", work);
		if (res != 0) {
			::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_UNPACK);
			MessageBoxA(this->GetHWND(), "请求拆包进度失败!", "警告", MB_OK | MB_ICONWARNING);
			return 0;
		}

		DWORD dwDone = StringToDword(done);
		DWORD dwWork = StringToDword(work) / 128;

		// update process information...
		int value = (int)((float)dwDone * 100 / (float)dwWork);
		if (value >= 100) {
			value = 100;
			::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_UNPACK);
		}
		m_pUnpackProgress->SetValue(value);

		break;
	}
	default:
		break;
	}

	return 0;
}

//----------------------------------------------
// @Function:	OnNcActivate()
// @Purpose: CFrameMain active message
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnNcActivate(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	if (::IsIconic(*this)) {
		bHandled = FALSE;
	}
	return (wParam == 0) ? TRUE : FALSE;
}

//----------------------------------------------
// @Function:	OnNcCalcSize()
// @Purpose: CFrameMain change size message
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnNcCalcSize(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	return 0;
}

//----------------------------------------------
// @Function:	OnNcPaint()
// @Purpose: CFrameMain paint message
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnNcPaint(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	return 0;
}

//----------------------------------------------
// @Function:	OnNcHitTest()
// @Purpose: CFrameMain hit message
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnNcHitTest(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	POINT pt;
	pt.x = GET_X_LPARAM(lParam);
	pt.y = GET_Y_LPARAM(lParam);
	::ScreenToClient(*this, &pt);

	RECT rcClient;
	::GetClientRect(*this, &rcClient);

	if (!::IsZoomed(*this)) {
		RECT rcSizeBox = m_PaintManager.GetSizeBox();
		if (pt.y < rcClient.top + rcSizeBox.top) {
			if (pt.x < rcClient.left + rcSizeBox.left) return HTTOPLEFT;
			if (pt.x > rcClient.right - rcSizeBox.right) return HTTOPRIGHT;
			return HTTOP;
		} else if (pt.y > rcClient.bottom - rcSizeBox.bottom) {
			if (pt.x < rcClient.left + rcSizeBox.left) return HTBOTTOMLEFT;
			if (pt.x > rcClient.right - rcSizeBox.right) return HTBOTTOMRIGHT;
			return HTBOTTOM;
		}
		if (pt.x < rcClient.left + rcSizeBox.left) return HTLEFT;
		if (pt.x > rcClient.right - rcSizeBox.right) return HTRIGHT;
	}

	RECT rcCaption = m_PaintManager.GetCaptionRect();
	if (pt.x >= rcClient.left + rcCaption.left && pt.x < rcClient.right - rcCaption.right && pt.y >= rcCaption.top && pt.y < rcCaption.bottom) {
		CControlUI* pControl = static_cast<CControlUI*>(m_PaintManager.FindControl(pt));
		if (pControl && _tcscmp(pControl->GetClass(), DUI_CTR_BUTTON) != 0 && _tcscmp(pControl->GetClass(), DUI_CTR_OPTION) != 0 && _tcscmp(pControl->GetClass(), DUI_CTR_TEXT) != 0) {
			return HTCAPTION;
		}
	}

	return HTCLIENT;
}

//----------------------------------------------
// @Function:	OnSize()
// @Purpose: CFrameMain change size message
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnSize(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	SIZE szRoundCorner = m_PaintManager.GetRoundCorner();

	if (!::IsIconic(*this) && (szRoundCorner.cx != 0 || szRoundCorner.cy != 0)) {
		CDuiRect rcWnd;
		::GetWindowRect(*this, &rcWnd);
		rcWnd.Offset(-rcWnd.left, -rcWnd.top);
		rcWnd.right++; rcWnd.bottom++;
		HRGN hRgn = ::CreateRoundRectRgn(rcWnd.left, rcWnd.top, rcWnd.right, rcWnd.bottom, szRoundCorner.cx, szRoundCorner.cy);
		::SetWindowRgn(*this, hRgn, TRUE);
		::DeleteObject(hRgn);
	}

	bHandled = FALSE;
	return 0;
}

//----------------------------------------------
// @Function:	OnGetMinMaxInfo()
// @Purpose: CFrameMain get max and min info
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnGetMinMaxInfo(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	MONITORINFO oMonitor = {};
	oMonitor.cbSize = sizeof(oMonitor);
	::GetMonitorInfo(::MonitorFromWindow(*this, MONITOR_DEFAULTTOPRIMARY), &oMonitor);
	CDuiRect rcWork = oMonitor.rcWork;
	rcWork.Offset(-oMonitor.rcMonitor.left, -oMonitor.rcMonitor.top);

	LPMINMAXINFO lpMMI = (LPMINMAXINFO)lParam;
	lpMMI->ptMaxPosition.x = rcWork.left;
	lpMMI->ptMaxPosition.y = rcWork.top;
	lpMMI->ptMaxSize.x = rcWork.right;
	lpMMI->ptMaxSize.y = rcWork.bottom;

	bHandled = FALSE;
	return 0;
}

//----------------------------------------------
// @Function:	OnSysCommand()
// @Purpose: CFrameMain system message
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnSysCommand(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	if (wParam == SC_CLOSE) {
		::PostQuitMessage(0L);
		bHandled = TRUE;
		return 0;
	}

	BOOL bZoomed = ::IsZoomed(*this);
	LRESULT lRes = CWindowWnd::HandleMessage(uMsg, wParam, lParam);
	if (::IsZoomed(*this) != bZoomed) {
		if (!bZoomed) {
			CControlUI* pControl = static_cast<CControlUI*>(m_PaintManager.FindControl(_T("maxbtn")));
			if (pControl) pControl->SetVisible(false);
			pControl = static_cast<CControlUI*>(m_PaintManager.FindControl(_T("restorebtn")));
			if (pControl) pControl->SetVisible(true);
		} else {
			CControlUI* pControl = static_cast<CControlUI*>(m_PaintManager.FindControl(_T("maxbtn")));
			if (pControl) pControl->SetVisible(true);
			pControl = static_cast<CControlUI*>(m_PaintManager.FindControl(_T("restorebtn")));
			if (pControl) pControl->SetVisible(false);
		}
	}
	return lRes;
}

//----------------------------------------------
// @Function:	OnUserMessagePacketSearch()
// @Purpose: CFrameMain packet search
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessagePacketSearch(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	HANDLE hThread = NULL;
	DWORD dwThreadID = 0;

	// clear content...
	m_pPackList->RemoveAll();
	m_pPackList->SetTextCallback(&g_cFramePackListUI);

	// create thread...
	hThread = ::CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)(&CFrameMain::OnSearchPacketItemsProcess), NULL, 0, &dwThreadID);
	::CloseHandle(hThread);

	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessagePacketAddItem()
// @Purpose: CFrameMain packet add items
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessagePacketAddItem(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	CListTextElementUI* pListElement = reinterpret_cast<CListTextElementUI*>(lParam);

	if (m_pPackList) {
		m_pPackList->Add(pListElement);
	}

	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessagePacketResult()
// @Purpose: CFrameMain packet display result
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessagePacketResult(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	if (lParam != 0) {
		::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_PACKET);
		MessageBoxA(this->GetHWND(), "请求打包文件失败!", "警告", MB_OK | MB_ICONWARNING);
	}
	else {
		MessageBoxA(this->GetHWND(), "打包文件成功!", "提示", MB_OK | MB_ICONASTERISK);
	}

	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessageUnpackSearch()
// @Purpose: CFrameMain unpack search
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessageUnpackSearch(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	HANDLE hThread = NULL;
	DWORD dwThreadID = 0;

	// clear content...
	m_pUnpackList->RemoveAll();
	m_pUnpackList->SetTextCallback(&g_cFrameUnpackListUI);

	// create thread...
	hThread = ::CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)(&CFrameMain::OnSearchUnpackItemsProcess), NULL, 0, &dwThreadID);
	::CloseHandle(hThread);

	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessageUnpackAddItem()
// @Purpose: CFrameMain unpack add items
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessageUnpackAddItem(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	CListTextElementUI* pListElement = reinterpret_cast<CListTextElementUI*>(lParam);

	if (m_pUnpackList) {
		m_pUnpackList->Add(pListElement);
	}

	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessageUnpackResult()
// @Purpose: CFrameMain unpack result
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessageUnpackResult(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	if (lParam != 0) {
		::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_UNPACK);
		MessageBoxA(this->GetHWND(), "请求拆包文件失败!", "警告", MB_OK | MB_ICONWARNING);
	}
	else {
		MessageBoxA(this->GetHWND(), "拆包文件成功!", "提示", MB_OK | MB_ICONASTERISK);
	}

	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessageCompSearch()
// @Purpose: CFrameMain compress search
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessageCompSearch(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	HANDLE hThread = NULL;
	DWORD dwThreadID = 0;

	// clear content...
	m_pCompList->RemoveAll();
	m_pCompList->SetTextCallback(&g_cFrameCompListUI);

	// create thread...
	hThread = ::CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)(&CFrameMain::OnSearchCompItemsProcess), NULL, 0, &dwThreadID);
	::CloseHandle(hThread);

	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessageCompAddItem()
// @Purpose: CFrameMain compress add item
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessageCompAddItem(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	CListTextElementUI* pListElement = reinterpret_cast<CListTextElementUI*>(lParam);

	if (m_pCompList) {
		m_pCompList->Add(pListElement);
	}

	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessageCompResult()
// @Purpose: CFrameMain compress result
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessageCompResult(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	if (lParam != 0) {
		::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_COMP);
		MessageBoxA(this->GetHWND(), "请求压缩文件失败!", "警告", MB_OK | MB_ICONWARNING);
	}
	else {
		MessageBoxA(this->GetHWND(), "打包压缩成功!", "提示", MB_OK | MB_ICONASTERISK);
	}

	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessageDecompSearch()
// @Purpose: CFrameMain compress search
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessageDecompSearch(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessageDecompAddItem()
// @Purpose: CFrameMain compress add item
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessageDecompAddItem(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	return 0;
}

//----------------------------------------------
// @Function:	OnUserMessageDecompResult()
// @Purpose: CFrameMain compress result
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessageDecompResult(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	if (lParam != 0) {
		::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_DECOMP);
		MessageBoxA(this->GetHWND(), "请求解压文件失败!", "警告", MB_OK | MB_ICONWARNING);
	}
	else {
		MessageBoxA(this->GetHWND(), "打包解压成功!", "提示", MB_OK | MB_ICONASTERISK);
	}

	return 0;
}

//----------------------------------------------
// @Function:	ConstructExtra()
// @Purpose: CFrameMain construct function extra
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::ConstructExtra() {
	m_hMenu = NULL;
	memset(&m_nid, 0, sizeof(m_nid));

	g_pFrameMain = this;
	m_vecPacket.clear();
	m_vecUnpack.clear();
	m_vecComp.clear();

	::srand((unsigned int)time(NULL));
}

//----------------------------------------------
// @Function:	DestructExtra()
// @Purpose: CFrameMain destruct function extra
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::DestructExtra() {
}

//----------------------------------------------
// @Function:	InitMenuShow()
// @Purpose: CFrameMain initialize menu show
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::InitMenuShow() {
	m_nid.cbSize = sizeof(NOTIFYICONDATA);
	m_nid.hWnd = this->GetHWND();
	m_nid.uID = IDI_ICON;
	m_nid.hIcon = ::LoadIcon(CPaintManagerUI::GetInstance(), MAKEINTRESOURCE(IDI_ICON));
	m_nid.uCallbackMessage = WM_USER_MESSAGE_MENU;
	m_nid.uFlags = NIF_ICON | NIF_MESSAGE | NIF_TIP | NIF_INFO;
	_tcscpy(m_nid.szTip, _T("detector"));
	Shell_NotifyIcon(NIM_ADD, &m_nid);

	m_hMenu = ::CreatePopupMenu();
	AppendMenu(m_hMenu, MF_STRING, ID_MAIN_EXIT, _T("Exit(Q)"));
}

//----------------------------------------------
// @Function:	InitWindowSharp()
// @Purpose: CFrameMain initialize window sharp
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::InitWindowSharp() {
	// change windows style ~ Areo
	::SetClassLongA(this->GetHWND(), GCL_STYLE, ::GetClassLongA(this->GetHWND(), GCL_STYLE) | CS_DROPSHADOW);
}

//----------------------------------------------
// @Function:	InitControls()
// @Purpose: CFrameMain initialize controls
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::InitControls() {
	// title buttons
	m_pCloseBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("closebtn")));
	m_pRestoreBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("restorebtn")));
	m_pMaxBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("maxbtn")));
	m_pMinBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("minbtn")));
	
	// menu buttons
	m_pMainTab = static_cast<CTabLayoutUI*>(m_PaintManager.FindControl(_T("maintab")));
	m_pPackOpt = static_cast<COptionUI*>(m_PaintManager.FindControl(_T("packopt")));
	m_pUnpackOpt = static_cast<COptionUI*>(m_PaintManager.FindControl(_T("unpackopt")));
	m_pCompOpt = static_cast<COptionUI*>(m_PaintManager.FindControl(_T("compopt")));
	m_pDecompOpt = static_cast<COptionUI*>(m_PaintManager.FindControl(_T("decompopt")));
	m_pBaseOpt = static_cast<COptionUI*>(m_PaintManager.FindControl(_T("baseopt")));
	m_pAppOpt = static_cast<COptionUI*>(m_PaintManager.FindControl(_T("appopt")));
	m_pMoreOpt = static_cast<COptionUI*>(m_PaintManager.FindControl(_T("moreopt")));
	m_pAboutOpt = static_cast<COptionUI*>(m_PaintManager.FindControl(_T("aboutopt")));

	// packet page
	m_pPackAddBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("packaddbtn")));
	m_pPackDelBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("packdelbtn")));
	m_pPackClrBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("packclrbtn")));
	m_pPackList = static_cast<CListUI*>(m_PaintManager.FindControl(_T("packlist")));
	m_pPackTypeEdt = static_cast<CEditUI*>(m_PaintManager.FindControl(_T("packtypeedt")));
	m_pPackPathEdt = static_cast<CEditUI*>(m_PaintManager.FindControl(_T("packpathedt")));
	m_pPackExportBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("packexportbtn")));
	m_pPackProgress = static_cast<CProgressUI*>(m_PaintManager.FindControl(_T("packprogressbar")));
	m_pPackStartBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("packstartbtn")));

	// unpack page
	m_pUnpackSelBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("unpackselbtn")));
	m_pUnpackUnselBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("unpackunselbtn")));
	m_pUnpackRetBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("unpackretbtn")));
	m_pUnpackUnretBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("unpackunretbtn")));
	m_pUnpackDetBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("unpackdetbtn")));
	m_pUnpackList = static_cast<CListUI*>(m_PaintManager.FindControl(_T("unpacklist")));
	m_pUnpackSrcEdt = static_cast<CEditUI*>(m_PaintManager.FindControl(_T("unpacksrcedt")));
	m_pUnpackDestEdt = static_cast<CEditUI*>(m_PaintManager.FindControl(_T("unpackdestedt")));
	m_pUnpackImportBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("unpackimportbtn")));
	m_pUnpackExportBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("unpackexportbtn")));
	m_pUnpackProgress = static_cast<CProgressUI*>(m_PaintManager.FindControl(_T("unpackprogressbar")));
	m_pUnpackStartBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("unpackstartbtn")));

	// compress page
	m_pCompAddBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("compaddbtn")));
	m_pCompDelBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("compdelbtn")));
	m_pCompClrBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("compclrbtn")));
	m_pCompList = static_cast<CListUI*>(m_PaintManager.FindControl(_T("complist")));
	m_pCompTypeEdt = static_cast<CEditUI*>(m_PaintManager.FindControl(_T("comptypeedt")));
	m_pCompPathEdt = static_cast<CEditUI*>(m_PaintManager.FindControl(_T("comppathedt")));
	m_pCompExportBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("compexportbtn")));
	m_pCompProgress = static_cast<CProgressUI*>(m_PaintManager.FindControl(_T("compprogressbar")));
	m_pCompStartBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("compstartbtn")));

	// decompress page
	m_pDecompDetBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("decompdetbtn")));
	m_pDecompList = static_cast<CListUI*>(m_PaintManager.FindControl(_T("decomplist")));
	m_pDecompSrcEdt = static_cast<CEditUI*>(m_PaintManager.FindControl(_T("decompsrcedt")));
	m_pDecompDestEdt = static_cast<CEditUI*>(m_PaintManager.FindControl(_T("decompdestedt")));
	m_pDecompImportBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("decompimportbtn")));
	m_pDecompExportBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("decompexportbtn")));
	m_pDecompProgress = static_cast<CProgressUI*>(m_PaintManager.FindControl(_T("decompprogressbar")));
	m_pDecompStartBtn = static_cast<CButtonUI*>(m_PaintManager.FindControl(_T("decompstartbtn")));
}

//----------------------------------------------
// @Function:	InitCurl()
// @Purpose: CFrameMain initialize libcurl
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::InitCurl() {
	CURLcode res = curl_global_init(CURL_GLOBAL_ALL);
	if (res != CURLE_OK) {
		MessageBoxA(this->GetHWND(), "初始化libcurl失败!", "错误", MB_OK | MB_ICONERROR);
		return;
	}
}

//----------------------------------------------
// @Function:	SplicePackRequestJson()
// @Purpose: CFrameMain splice packet json
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CDuiString CFrameMain::SplicePackRequestJson(CDuiString strPacketType, CDuiString strPacketPath) {
	CDuiString strPacketJson;
	CDuiString strPacketSrc;
	CDuiString strPacketDest;
	CDuiString strPacketKind;

	USES_CONVERSION;

	// splice src files...
	strPacketSrc = L"\"src\":[";
	for (auto iter = m_vecPacket.begin(); iter != m_vecPacket.end(); ++iter) {
		char chFile[MAX_PATH] = { 0 };
		if (iter == m_vecPacket.end() - 1) {
			sprintf_s(chFile, "\"%s\"", iter->chPath);
		} else {
			sprintf_s(chFile, "\"%s\",", iter->chPath);
		}
		strPacketSrc += A2T(chFile);
	}
	strPacketSrc += L"],";
	//strPacketSrc.Replace(L"\\", L"/");

	// splice dest file...
	strPacketDest = L"\"dest\":\"";
	strPacketDest += strPacketPath;
	strPacketDest += L"\",";

	// splice encrypt type...
	strPacketKind = L"\"type\":\"";
	strPacketKind += strPacketType;
	strPacketKind += L"\"";

	// splice all...
	strPacketJson = L"{";
	strPacketJson += strPacketSrc;
	strPacketJson += strPacketDest;
	strPacketJson += strPacketKind;
	strPacketJson += L"}";
	strPacketJson.Replace(L"\\", L"/");

	return strPacketJson;
}

//----------------------------------------------
// @Function:	SplicePackProcessRequestJson()
// @Purpose: CFrameMain splice packet process json
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CDuiString CFrameMain::SplicePackProcessRequestJson(CDuiString strPacketType) {
	CDuiString strPacketJson;
	CDuiString strPacketSrc;
	CDuiString strPacketDest;
	CDuiString strPacketKind;

	USES_CONVERSION;

	// splice src files...
	strPacketSrc = L"\"src\":[";
	for (auto iter = m_vecPacket.begin(); iter != m_vecPacket.end(); ++iter) {
		char chFile[MAX_PATH] = { 0 };
		if (iter == m_vecPacket.end() - 1) {
			sprintf_s(chFile, "\"%s\"", iter->chPath);
		}
		else {
			sprintf_s(chFile, "\"%s\",", iter->chPath);
		}
		strPacketSrc += A2T(chFile);
	}
	strPacketSrc += L"],";
	//strPacketSrc.Replace(L"\\", L"/");

	// splice encrypt type...
	strPacketKind = L"\"type\":\"";
	strPacketKind += strPacketType;
	strPacketKind += L"\"";

	// splice all...
	strPacketJson = L"{";
	strPacketJson += strPacketSrc;
	strPacketJson += strPacketKind;
	strPacketJson += L"}";
	strPacketJson.Replace(L"\\", L"/");

	return strPacketJson;
}

//----------------------------------------------
// @Function:	SplicePackProcessRequestJson()
// @Purpose: CFrameMain splice unpack process json
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CDuiString CFrameMain::SpliceUnpackRequestJson(CDuiString strUnpackSrc, CDuiString strUnpackDest) {
	CDuiString strUnpackJson;
	CDuiString strUnpackSource;
	CDuiString strUnpackDestination;

	USES_CONVERSION;

	// splice src files...
	strUnpackSource = L"\"src\":";
	strUnpackSource += L"\"";
	strUnpackSource += strUnpackSrc.GetData();
	strUnpackSource += L"\",";

	// splice dest files...
	strUnpackDestination = L"\"dest\":";
	strUnpackDestination += L"\"";
	strUnpackDestination += strUnpackDest.GetData();
	strUnpackDestination += L"\"";

	// splice all...
	strUnpackJson = L"{";
	strUnpackJson += strUnpackSource;
	strUnpackJson += strUnpackDestination;
	strUnpackJson += L"}";
	strUnpackJson.Replace(L"\\", L"/");

	return strUnpackJson;
}

//----------------------------------------------
// @Function:	SpliceUnpackVerboseRequestJson()
// @Purpose: CFrameMain splice unpack verbose process json
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CDuiString CFrameMain::SpliceUnpackVerboseRequestJson(CDuiString strUnpackSrc) {
	CDuiString strUnpackJson;
	CDuiString strUnpackSource;

	USES_CONVERSION;

	// splice src files...
	strUnpackSource = L"\"src\":";
	strUnpackSource += L"\"";
	strUnpackSource += strUnpackSrc.GetData();
	strUnpackSource += L"\"";

	// splice all...
	strUnpackJson = L"{";
	strUnpackJson += strUnpackSource;
	strUnpackJson += L"}";
	strUnpackJson.Replace(L"\\", L"/");

	return strUnpackJson;
}

//----------------------------------------------
// @Function:	SpliceUnpackProcessRequestJson()
// @Purpose: CFrameMain splice unpack process process json
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CDuiString CFrameMain::SpliceUnpackProcessRequestJson(CDuiString strUnpackSrc) {
	CDuiString strUnpackJson;
	CDuiString strUnpackSource;

	USES_CONVERSION;

	// splice src files...
	strUnpackSource = L"\"src\":";
	strUnpackSource += L"\"";
	strUnpackSource += strUnpackSrc.GetData();
	strUnpackSource += L"\"";

	// splice all...
	strUnpackJson = L"{";
	strUnpackJson += strUnpackSource;
	strUnpackJson += L"}";
	strUnpackJson.Replace(L"\\", L"/");

	return strUnpackJson;
}

//----------------------------------------------
// @Function:	SpliceUnpackOneFileRequestJson()
// @Purpose: CFrameMain splice unpack one file process json
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CDuiString CFrameMain::SpliceUnpackOneFileRequestJson(CDuiString strUnpackSrc, CDuiString strUnpackDest, CDuiString strUnpackFile) {
	CDuiString strUnpackJson;
	CDuiString strUnpackSource;
	CDuiString strUnpackDestination;
	CDuiString strUnpackTarget;

	USES_CONVERSION;

	// splice src files...
	strUnpackSource = L"\"src\":";
	strUnpackSource += L"\"";
	strUnpackSource += strUnpackSrc.GetData();
	strUnpackSource += L"\",";

	// splice target files...
	strUnpackTarget = L"\"target\":";
	strUnpackTarget += L"\"";
	strUnpackTarget += strUnpackFile.GetData();
	strUnpackTarget += L"\",";

	// splice dest files...
	strUnpackDestination = L"\"dest\":";
	strUnpackDestination += L"\"";
	strUnpackDestination += strUnpackDest.GetData();
	strUnpackDestination += L"\"";

	// splice all...
	strUnpackJson = L"{";
	strUnpackJson += strUnpackSource;
	strUnpackJson += strUnpackTarget;
	strUnpackJson += strUnpackDestination;
	strUnpackJson += L"}";
	strUnpackJson.Replace(L"\\", L"/");

	return strUnpackJson;
}

//----------------------------------------------
// @Function:	SpliceCompRequestJson()
// @Purpose: CFrameMain compress request json
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CDuiString CFrameMain::SpliceCompRequestJson(CDuiString strCompType, CDuiString strCompPath) {
	CDuiString strCompJson;
	CDuiString strCompSrc;
	CDuiString strCompDest;
	CDuiString strCompKind;

	USES_CONVERSION;

	// splice src files...
	strCompSrc = L"\"src\":[";
	for (auto iter = m_vecComp.begin(); iter != m_vecComp.end(); ++iter) {
		char chFile[MAX_PATH] = { 0 };
		if (iter == m_vecComp.end() - 1) {
			sprintf_s(chFile, "\"%s\"", iter->chPath);
		}
		else {
			sprintf_s(chFile, "\"%s\",", iter->chPath);
		}
		strCompSrc += A2T(chFile);
	}
	strCompSrc += L"],";

	// splice dest file...
	strCompDest = L"\"dest\":\"";
	strCompDest += strCompPath;
	strCompDest += L"\",";

	// splice encrypt type...
	strCompKind = L"\"type\":\"";
	strCompKind += strCompType;
	strCompKind += L"\"";

	// splice all...
	strCompJson = L"{";
	strCompJson += strCompSrc;
	strCompJson += strCompDest;
	strCompJson += strCompKind;
	strCompJson += L"}";
	strCompJson.Replace(L"\\", L"/");

	return strCompJson;
}

//----------------------------------------------
// @Function:	SpliceDecompRequestJson()
// @Purpose: CFrameMain compress request json
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CDuiString CFrameMain::SpliceDecompRequestJson(CDuiString strDecompSrc, CDuiString strDecompDest, CDuiString strDecompType) {
	CDuiString strDecompJson;
	CDuiString strDecompSource;
	CDuiString strDecompDestination;
	CDuiString strCompKind;

	USES_CONVERSION;

	// splice src files...
	strDecompSource = L"\"src\":";
	strDecompSource += L"\"";
	strDecompSource += strDecompSrc.GetData();
	strDecompSource += L"\",";

	// splice dest files...
	strDecompDestination = L"\"dest\":";
	strDecompDestination += L"\"";
	strDecompDestination += strDecompDest.GetData();
	strDecompDestination += L"\",";

	// splice encrypt type...
	strCompKind = L"\"type\":\"";
	strCompKind += strDecompType.GetData();
	strCompKind += L"\"";

	// splice all...
	strDecompJson = L"{";
	strDecompJson += strDecompSource;
	strDecompJson += strDecompDestination;
	strDecompJson += strCompKind;
	strDecompJson += L"}";
	strDecompJson.Replace(L"\\", L"/");

	return strDecompJson;
}

//----------------------------------------------
// @Function:	GetValueFromResponse()
// @Purpose: CFrameMain get value from response
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
int CFrameMain::GetValueFromResponse(string result, string label_first, string label_last, string& value) {
	int size;
	int pos;
	int end;
	int count;
	size = label_first.length();
	pos = result.find(label_first);
	if (pos == result.npos) {
		return -1;
	}
	end = result.find_first_of(label_last, pos);
	if (pos == result.npos) {
		return -1;
	}
	count = end - pos - size;
	value = result.substr(pos + size, count);
	return 0;
}

//----------------------------------------------
// @Function:	GetValueFromResponseWithPos()
// @Purpose: CFrameMain get value from response
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
int CFrameMain::GetValueFromResponseWithPos(string result, string label_first, string label_last, string& value, int& poi) {
	int size;
	int pos;
	int end;
	int count;
	size = label_first.length();
	pos = result.find(label_first);
	if (pos == result.npos) {
		return -1;
	}
	end = result.find_first_of(label_last, pos);
	if (pos == result.npos) {
		return -1;
	}
	count = end - pos - size;
	value = result.substr(pos + size, count);
	poi = end;
	return 0;
}

//----------------------------------------------
// @Function:	OnSearchPacketItemsProcess()
// @Purpose: CFrameMain search packet items
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::OnSearchPacketItemsProcess(LPVOID lpParameter) {
	CListUI* pList = g_pFrameMain->m_pPackList;

	for (int i = 0; i < g_pFrameMain->m_vecPacket.size(); ++i) {
		CListTextElementUI* pListElement = new CListTextElementUI();

		pListElement->SetTag(i);
		if (pListElement != NULL) {
			::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_PACKET_ADDITEM, 0L, (LPARAM)pListElement);
		}

		Sleep(1);
	}

	return 0;
}

//----------------------------------------------
// @Function:	OnSearchUnpackItemsProcess()
// @Purpose: CFrameMain search unpack items
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::OnSearchUnpackItemsProcess(LPVOID lpParameter) {
	CListUI* pList = g_pFrameMain->m_pUnpackList;

	for (int i = 0; i < g_pFrameMain->m_vecUnpack.size(); ++i) {
		CListTextElementUI* pListElement = new CListTextElementUI();

		pListElement->SetTag(i);
		if (pListElement != NULL) {
			::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_UNPACK_ADDITEM, 0L, (LPARAM)pListElement);
		}

		Sleep(1);
	}

	return 0;
}

//----------------------------------------------
// @Function:	OnSearchCompItemsProcess()
// @Purpose: CFrameMain search compress items
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::OnSearchCompItemsProcess(LPVOID lpParameter) {
	CListUI* pList = g_pFrameMain->m_pCompList;

	for (int i = 0; i < g_pFrameMain->m_vecComp.size(); ++i) {
		CListTextElementUI* pListElement = new CListTextElementUI();

		pListElement->SetTag(i);
		if (pListElement != NULL) {
			::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_COMP_ADDITEM, 0L, (LPARAM)pListElement);
		}

		Sleep(1);
	}

	return 0;
}

//----------------------------------------------
// @Function:	OnPostPackRequestProcess()
// @Purpose: CFrameMain post pack request
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::OnPostPackRequestProcess(LPVOID lpParameter) {
	string data = string((char*)lpParameter);
	string url;
	string response;
	CURLcode res;

	url = "http://localhost:8080/satellite/pack";
	res = CurlPostRequest(url, data, response);
	if (res != CURLE_OK) {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_PACKET_RESULT, (WPARAM)0, (LPARAM)1);
	} else {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_PACKET_RESULT, (WPARAM)0, (LPARAM)0);
	}
	return 0;
}

//----------------------------------------------
// @Function:	OnPostUnpackRequestProcess()
// @Purpose: CFrameMain post unpack request
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::OnPostUnpackRequestProcess(LPVOID lpParameter) {
	string data = string((char*)lpParameter);
	string url;
	string response;
	CURLcode res;

	url = "http://localhost:8080/satellite/unpack";
	res = CurlPostRequest(url, data, response);
	if (res != CURLE_OK) {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_UNPACK_RESULT, (WPARAM)0, (LPARAM)1);
	}
	else {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_UNPACK_RESULT, (WPARAM)0, (LPARAM)0);
	}
	return 0;
}

//----------------------------------------------
// @Function:	OnPostUnpackConfineRequestProcess()
// @Purpose: CFrameMain post unpack confine request
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::OnPostUnpackConfineRequestProcess(LPVOID lpParameter) {
	string data = string((char*)lpParameter);
	string url;
	string response;
	CURLcode res;

	url = "http://localhost:8080/satellite/unpack/c";
	res = CurlPostRequest(url, data, response);
	if (res != CURLE_OK) {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_UNPACK_RESULT, (WPARAM)0, (LPARAM)1);
	}
	else {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_UNPACK_RESULT, (WPARAM)0, (LPARAM)0);
	}
	return 0;
}

//----------------------------------------------
// @Function:	OnPostUnpackFileRequestProcess()
// @Purpose: CFrameMain post unpack file request
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::OnPostUnpackFileRequestProcess(LPVOID lpParameter) {
	string data = string((char*)lpParameter);
	string url;
	string response;
	CURLcode res;

	url = "http://localhost:8080/satellite/unpack/f";
	res = CurlPostRequest(url, data, response);
	if (res != CURLE_OK) {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_UNPACK_RESULT, (WPARAM)0, (LPARAM)1);
	}
	else {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_UNPACK_RESULT, (WPARAM)0, (LPARAM)0);
	}
	return 0;
}

//----------------------------------------------
// @Function:	OnPostUnpackFileConfineRequestProcess()
// @Purpose: CFrameMain post unpack file confine request
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::OnPostUnpackFileConfineRequestProcess(LPVOID lpParameter) {
	string data = string((char*)lpParameter);
	string url;
	string response;
	CURLcode res;

	url = "http://localhost:8080/satellite/unpack/cf";
	res = CurlPostRequest(url, data, response);
	if (res != CURLE_OK) {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_UNPACK_RESULT, (WPARAM)0, (LPARAM)1);
	}
	else {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_UNPACK_RESULT, (WPARAM)0, (LPARAM)0);
	}
	return 0;
}

//----------------------------------------------
// @Function:	OnPostCompRequestProcess()
// @Purpose: CFrameMain post compress request
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::OnPostCompRequestProcess(LPVOID lpParameter) {
	string data = string((char*)lpParameter);
	string url;
	string response;
	CURLcode res;

	url = "http://localhost:8080/satellite/comp";
	res = CurlPostRequest(url, data, response);
	if (res != CURLE_OK) {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_COMP_RESULT, (WPARAM)0, (LPARAM)1);
	}
	else {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_COMP_RESULT, (WPARAM)0, (LPARAM)0);
	}
	return 0;
}

//----------------------------------------------
// @Function:	OnPostDecompRequestProcess()
// @Purpose: CFrameMain post decompress request
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::OnPostDecompRequestProcess(LPVOID lpParameter) {
	string data = string((char*)lpParameter);
	string url;
	string response;
	CURLcode res;

	url = "http://localhost:8080/satellite/decomp";
	res = CurlPostRequest(url, data, response);
	if (res != CURLE_OK) {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_DECOMP_RESULT, (WPARAM)0, (LPARAM)1);
	}
	else {
		::PostMessageA(g_pFrameMain->GetHWND(), WM_USER_MESSAGE_DECOMP_RESULT, (WPARAM)0, (LPARAM)0);
	}
	return 0;
}

//----------------------------------------------
// @Function:	StringToDword()
// @Purpose: CFrameMain convert between string & dword
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
DWORD CFrameMain::StringToDword(string value) {
	DWORD dword;
	sscanf(value.c_str(), "%ul", &dword);
	return dword;
}

//----------------------------------------------
// @Function:	StringSplit()
// @Purpose: CFrameMain split string
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::StringSplit(const std::string& s, std::vector<std::string>& v, const std::string& c) {
	string::size_type pos1, pos2;
	pos2 = s.find(c);
	pos1 = 0;
	while(string::npos != pos2) {
		v.push_back(s.substr(pos1, pos2-pos1));

		pos1 = pos2 + c.size();
		pos2 = s.find(c, pos1);
	}
	if (pos1 != s.length()) {
		v.push_back(s.substr(pos1));
	}
}

//----------------------------------------------
// @Function:	OnUserMessageMenu()
// @Purpose: CFrameMain menu user message
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
LRESULT CFrameMain::OnUserMessageMenu(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled) {
	switch (lParam) {
	case WM_RBUTTONDOWN: {
		POINT pt;
		int nRet;

		GetCursorPos(&pt);
		::SetForegroundWindow(this->GetHWND());

		nRet = TrackPopupMenu(m_hMenu, TPM_RETURNCMD, pt.x, pt.y, NULL, this->GetHWND(), NULL);
		if (nRet == ID_MAIN_EXIT) {
			::PostMessageA(this->GetHWND(), WM_CLOSE, (WPARAM)0, (LPARAM)0);
		}
	}
	break;
	case WM_LBUTTONDOWN:
	case WM_LBUTTONDBLCLK: {
		::ShowWindow(this->GetHWND(), SW_SHOW);
	}
	break;
	default:
		break;
	}

	return 0;
}

//----------------------------------------------
// @Function:	GetPaintManager()
// @Purpose: CFrameMain get paint manager handle
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
CPaintManagerUI& CFrameMain::GetPaintManager() {
	return m_PaintManager;
}

//----------------------------------------------
// @Function:	OnLButtonClickedMinBtn()
// @Purpose: CFrameMain click min button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedMinBtn() {
	SendMessage(WM_SYSCOMMAND, SC_MINIMIZE, 0);
}

//----------------------------------------------
// @Function:	OnLButtonClickedMaxBtn()
// @Purpose: CFrameMain click max button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedMaxBtn() {
	SendMessage(WM_SYSCOMMAND, SC_MAXIMIZE, 0);
}

//----------------------------------------------
// @Function:	OnLButtonClickedRestoreBtn()
// @Purpose: CFrameMain click restore button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedRestoreBtn() {
	SendMessage(WM_SYSCOMMAND, SC_RESTORE, 0);
}

//----------------------------------------------
// @Function:	OnLButtonClickedCloseBtn()
// @Purpose: CFrameMain click close button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedCloseBtn() {
	::ShowWindow(this->GetHWND(), SW_HIDE);
}

//----------------------------------------------
// @Function:	OnLButtonClickedCloseBtn()
// @Purpose: CFrameMain click packet add button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedPacketAddBtn() {
	OPENFILENAME file;
	WCHAR strfile[100 * MAX_PATH] = { 0 };
	WCHAR strpath[MAX_PATH] = { 0 };
	WCHAR strname[MAX_PATH] = { 0 };
	TCHAR* p = NULL;
	int nLen = 0;

	USES_CONVERSION;
	ZeroMemory(&file, sizeof(OPENFILENAME));
	file.lStructSize = sizeof(OPENFILENAME);
	file.lpstrFilter = _T("所有文件\0*.*\0\0");
	file.nFilterIndex = 1;
	file.lpstrFile = strfile;
	file.nMaxFile = sizeof(strfile);
	file.Flags = OFN_EXPLORER | OFN_ALLOWMULTISELECT;

	if (GetOpenFileName(&file)) {
		lstrcpyn(strpath, strfile, file.nFileOffset);
		strpath[file.nFileOffset] = '\0';
		nLen = lstrlen(strpath);

		if (strpath[nLen - 1] != '\\') {
			lstrcat(strpath, _T("\\"));
		}

		p = strfile + file.nFileOffset;
		while (*p) {
			ZeroMemory(strname, sizeof(strname));
			lstrcat(strname, strpath);
			lstrcat(strname, p);

			char chOriginFile[MAX_PATH] = { 0 };
			char chOriginName[MAX_PATH] = { 0 };
			char* pTemp = NULL;

			strcpy_s(chOriginFile, T2A(strname));
			pTemp = strrchr(chOriginFile, '\\');
			strcpy_s(chOriginName, ++pTemp);

			bool bRepeat = false;
			for (auto iter = m_vecPacket.begin(); iter != m_vecPacket.end(); ++iter) {
				if (!strcmp(iter->chPath, T2A(strname))) {
					bRepeat = true;
					break;
				}
			}

			if (!bRepeat) {
				TPacketInfo sPacketInfo = { 0 };
				sPacketInfo.nSerial = m_vecPacket.size() + 1;
				strcpy_s(sPacketInfo.chName, chOriginName);
				strcpy_s(sPacketInfo.chPath, T2A(strname));
				m_vecPacket.push_back(sPacketInfo);
			}
			p += lstrlen(p) + 1;
		}

		::PostMessageA(this->GetHWND(), WM_USER_MESSAGE_PACKET_SEARCH, (WPARAM)0, (LPARAM)0);
	}
}

//----------------------------------------------
// @Function:	OnLButtonClickedCloseBtn()
// @Purpose: CFrameMain click packet delete button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedPacketDelBtn() {
	int nItem = m_pPackList->GetCurSel();
	if (nItem < 0) {
		MessageBoxW(this->GetHWND(), _T("请选中一条文件信息!"), _T("提示"), MB_OK | MB_ICONASTERISK);
		return;
	}

	m_vecPacket.erase(m_vecPacket.begin() + nItem);
	for (auto iter = m_vecPacket.begin(); iter != m_vecPacket.end(); ++iter) {
		iter->nSerial = iter - m_vecPacket.begin() + 1;
	}

	::PostMessageA(this->GetHWND(), WM_USER_MESSAGE_PACKET_SEARCH, (WPARAM)0, (LPARAM)0);
}

//----------------------------------------------
// @Function:	OnLButtonClickedPacketClrBtn()
// @Purpose: CFrameMain click packet clear button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedPacketClrBtn() {
	m_vecPacket.clear();

	::PostMessageA(this->GetHWND(), WM_USER_MESSAGE_PACKET_SEARCH, (WPARAM)0, (LPARAM)0);
}

//----------------------------------------------
// @Function:	OnLButtonClickedPacketExportBtn()
// @Purpose: CFrameMain click packet export button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedPacketExportBtn() {
	OPENFILENAME file;
	WCHAR strfile[MAX_PATH] = { 0 };

	USES_CONVERSION;
	wcscpy_s(strfile, A2T("undefine.pak"));

	ZeroMemory(&file, sizeof(OPENFILENAME));
	file.lStructSize = sizeof(OPENFILENAME);
	file.lpstrFilter = _T("所有文件\0*.*\0\0");
	file.nFilterIndex = 1;
	file.lpstrFile = strfile;
	file.nMaxFile = sizeof(strfile);
	file.Flags = OFN_PATHMUSTEXIST | OFN_OVERWRITEPROMPT;

	if (GetSaveFileName(&file)) {
		m_pPackPathEdt->SetText(strfile);
	}
}

//----------------------------------------------
// @Function:	OnLButtonClickedPacketStartBtn()
// @Purpose: CFrameMain click packet start button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedPacketStartBtn() {
	// check import files...
	if (m_vecPacket.empty()) {
		MessageBoxA(this->GetHWND(), "请添加至少一个封包文件!", "提示", MB_OK | MB_ICONASTERISK);
		return;
	}

	USES_CONVERSION;

	// check encrypt type...
	CDuiString strPacketType = m_pPackTypeEdt->GetText();
	string strTypes[] = { "AES", "DES", "3DES", "RSA", "BASE64" };
	vector<string> vecPacketTypes(strTypes, strTypes + 5);
	bool bType = false;

	for (auto iter = vecPacketTypes.begin(); iter != vecPacketTypes.end(); ++iter) {
		if (!strcmp(iter->c_str(), T2A(strPacketType.GetData()))) {
			bType = true;
			break;
		}
	}

	if (bType == false) {
		MessageBoxA(this->GetHWND(), "请填入支持的加密类型(AES, DES, 3DES, RSA, BASE64)!", "提示", MB_OK | MB_ICONASTERISK);
		return;
	}

	// check export package...
	CDuiString strPacketPath = m_pPackPathEdt->GetText();
	if (!strcmp("", T2A(strPacketPath.GetData()))) {
		MessageBoxA(this->GetHWND(), "导出文件路径不能为空!", "警告", MB_OK | MB_ICONWARNING);
		return;
	}

	// set progress...
	m_pPackProgress->SetMinValue(0);
	m_pPackProgress->SetMaxValue(100);
	m_pPackProgress->SetValue(0);

	// reset timer...
	::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_PACKET);
	::SetTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_PACKET, CONST_PROGRESS_REFRESH_TIME, NULL);

	// organize http request...
	CDuiString strPacketJson;
	strPacketJson = SplicePackRequestJson(strPacketType, strPacketPath);

	// create thread handle request...
	HANDLE hThread = NULL;
	DWORD dwThreadID = 0;

	// create thread...
	hThread = ::CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)(&CFrameMain::OnPostPackRequestProcess), (LPVOID)(T2A(strPacketJson.GetData())), 0, &dwThreadID);
	::CloseHandle(hThread);

}

//----------------------------------------------
// @Function:	OnLButtonClickedUnpackSelectBtn()
// @Purpose: CFrameMain click unpack select button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedUnpackSelectBtn() {
	m_pUnpackSelBtn->SetVisible(false);
	m_pUnpackUnselBtn->SetVisible(true);
}

//----------------------------------------------
// @Function:	OnLButtonClickedUnpackUnselectBtn()
// @Purpose: CFrameMain click unpack unselect button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedUnpackUnselectBtn() {
	m_pUnpackSelBtn->SetVisible(true);
	m_pUnpackUnselBtn->SetVisible(false);
}

//----------------------------------------------
// @Function:	OnLButtonClickedUnpackRestrictBtn()
// @Purpose: CFrameMain click unpack restrict button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedUnpackRestrictBtn() {
	m_pUnpackRetBtn->SetVisible(false);
	m_pUnpackUnretBtn->SetVisible(true);
}

//----------------------------------------------
// @Function:	OnLButtonClickedUnpackUnrestrictBtn()
// @Purpose: CFrameMain click unpack restrict button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedUnpackUnrestrictBtn() {
	m_pUnpackRetBtn->SetVisible(true);
	m_pUnpackUnretBtn->SetVisible(false);
}

//----------------------------------------------
// @Function:	OnLButtonClickedUnpackDetialBtn()
// @Purpose: CFrameMain click unpack detial button
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedUnpackDetialBtn() {
	CDuiString strUnpackJson;
	CDuiString strUnpackSrc;

	USES_CONVERSION;

	strUnpackSrc = m_pUnpackSrcEdt->GetText();

	// check source file...
	if (strUnpackSrc.IsEmpty()) {
		MessageBoxA(this->GetHWND(), "源文件路径不能为空!", "警告", MB_OK | MB_ICONWARNING);
		return;
	}

	// organize http request...
	strUnpackJson = SpliceUnpackVerboseRequestJson(strUnpackSrc);

	// post http request...
	string url;
	string response;
	CURLcode res;

	url = "http://localhost:8080/satellite/unpack/v";
	res = CurlPostRequest(url, string(T2A(strUnpackJson.GetData())), response);
	if (res != CURLE_OK) {
		MessageBoxA(this->GetHWND(), "请求拆包文件详细信息失败!", "警告", MB_OK | MB_ICONWARNING);
		return;
	}

	// withdraw data from response...
	vector<string> v;
	
	// split strings
	StringSplit(response, v, "},");
	m_vecUnpack.clear();
	for (auto iter = v.begin(); iter != v.end(); ++iter) {
		string name;
		string size;
		int pos;
		int res;
		
		// get file name...
		res = GetValueFromResponseWithPos(*iter, "\"name\": \"", ",", name, pos);
		if (res != 0) {
			continue;
		}

		// get file size...
		res = GetValueFromResponse(string((*iter).c_str() + pos), "\"size\": \"", ",", size);
		if (res != 0) {
			continue;
		}

		name.erase(name.length() - 1);
		size.erase(size.length() - 1);

		TUnpackInfo sUnpackInfo = { 0 };
		sUnpackInfo.nSerial = m_vecUnpack.size() + 1;
		strcpy_s(sUnpackInfo.chName, name.c_str());
		strcpy_s(sUnpackInfo.chSize, size.c_str());
		m_vecUnpack.push_back(sUnpackInfo);
	}

	::PostMessageA(this->GetHWND(), WM_USER_MESSAGE_UNPACK_SEARCH, (WPARAM)0, (LPARAM)0);
}

//----------------------------------------------
// @Function:	OnLButtonClickedUnpackImportBtn()
// @Purpose: CFrameMain click unpack import
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedUnpackImportBtn() {
	OPENFILENAME file;
	WCHAR strfile[MAX_PATH] = { 0 };

	ZeroMemory(&file, sizeof(OPENFILENAME));
	file.lStructSize = sizeof(OPENFILENAME);
	file.lpstrFilter = _T("所有文件\0*.*\0\0");
	file.nFilterIndex = 1;
	file.lpstrFile = strfile;
	file.nMaxFile = sizeof(strfile);
	file.Flags = OFN_FILEMUSTEXIST | OFN_PATHMUSTEXIST;

	if (GetOpenFileName(&file)) {
		m_pUnpackSrcEdt->SetText(strfile);
	}
}

//----------------------------------------------
// @Function:	OnLButtonClickedUnpackExportBtn()
// @Purpose: CFrameMain click unpack export
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedUnpackExportBtn() {
	CDuiString strImport = _T("");
	WCHAR strfile[MAX_PATH] = { 0 };

	USES_CONVERSION;

	wcscpy_s(strfile, A2T("undefine.pak"));

	strImport = m_pUnpackSrcEdt->GetText();
	if (!strcmp(T2A(strImport.GetData()), "")) {
		MessageBoxA(this->GetHWND(), "请选择源文件路径!", "提示", MB_OK | MB_ICONASTERISK);
		return;
	}

	OPENFILENAME file;

	ZeroMemory(&file, sizeof(OPENFILENAME));
	file.lStructSize = sizeof(OPENFILENAME);
	file.lpstrFilter = _T("所有文件\0*.*\0\0");
	file.nFilterIndex = 1;
	file.lpstrFile = strfile;
	file.nMaxFile = sizeof(strfile);
	file.Flags = OFN_PATHMUSTEXIST | OFN_OVERWRITEPROMPT;

	if (GetSaveFileName(&file)) {
		CDuiString dest(strfile);
		dest = dest.Left(dest.ReverseFind(L'\\') + 1);
		m_pUnpackDestEdt->SetText(dest.GetData());
	}
}

//----------------------------------------------
// @Function:	OnLButtonClickedUnpackStartBtn()
// @Purpose: CFrameMain click unpack start
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedUnpackStartBtn() {
	CDuiString strUnpackJson;
	CDuiString strUnpackSrc;
	CDuiString strUnpackDest;
	CDuiString strUnpackFile;

	USES_CONVERSION;

	strUnpackSrc = m_pUnpackSrcEdt->GetText();
	strUnpackDest = m_pUnpackDestEdt->GetText();

	// check source file...
	if (strUnpackSrc.IsEmpty()) {
		MessageBoxA(this->GetHWND(), "源文件路径不能为空!", "警告", MB_OK | MB_ICONWARNING);
		return;
	}

	// check destination file...
	if (strUnpackDest.IsEmpty()) {
		MessageBoxA(this->GetHWND(), "目标文件路径不能为空!", "警告", MB_OK | MB_ICONWARNING);
		return;
	}

	if (!m_pUnpackSelBtn->IsVisible()) {
		// get select item...
		int nItem = m_pUnpackList->GetCurSel();
		if (nItem < 0 || nItem >= m_vecUnpack.size()) {
			MessageBoxW(this->GetHWND(), _T("请选中一条文件信息!"), _T("提示"), MB_OK | MB_ICONASTERISK);
			return;
		}

		// get file name...
		strUnpackFile = A2T(m_vecUnpack.at(nItem).chName);

		// check target file...
		if (strUnpackFile.IsEmpty()) {
			MessageBoxA(this->GetHWND(), "选中文件路径不能为空!", "警告", MB_OK | MB_ICONWARNING);
			return;
		}

		// organize http request...
		strUnpackJson = SpliceUnpackOneFileRequestJson(strUnpackSrc, strUnpackDest, strUnpackFile);
	} else {
		// organize http request...
		strUnpackJson = SpliceUnpackRequestJson(strUnpackSrc, strUnpackDest);
	}

	// set progress...
	m_pUnpackProgress->SetMinValue(0);
	m_pUnpackProgress->SetMaxValue(100);
	m_pUnpackProgress->SetValue(0);

	// reset timer...
	::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_UNPACK);
	::SetTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_UNPACK, CONST_PROGRESS_REFRESH_TIME, NULL);

	// create thread handle request...
	HANDLE hThread = NULL;
	DWORD dwThreadID = 0;

	// create thread...
	if (m_pUnpackSelBtn->IsVisible()) {
		if (m_pUnpackRetBtn->IsVisible()) {
			hThread = ::CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)(&CFrameMain::OnPostUnpackRequestProcess), (LPVOID)(T2A(strUnpackJson.GetData())), 0, &dwThreadID);
			::CloseHandle(hThread);
		} else {
			hThread = ::CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)(&CFrameMain::OnPostUnpackConfineRequestProcess), (LPVOID)(T2A(strUnpackJson.GetData())), 0, &dwThreadID);
			::CloseHandle(hThread);
		}
	} else {
		if (m_pUnpackRetBtn->IsVisible()) {
			hThread = ::CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)(&CFrameMain::OnPostUnpackFileRequestProcess), (LPVOID)(T2A(strUnpackJson.GetData())), 0, &dwThreadID);
			::CloseHandle(hThread);
		} else {
			hThread = ::CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)(&CFrameMain::OnPostUnpackFileConfineRequestProcess), (LPVOID)(T2A(strUnpackJson.GetData())), 0, &dwThreadID);
			::CloseHandle(hThread);
		}
	}
	

}

//----------------------------------------------
// @Function:	OnLButtonClickedCompAddBtn()
// @Purpose: CFrameMain click compress add
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedCompAddBtn() {
	OPENFILENAME file;
	WCHAR strfile[100 * MAX_PATH] = { 0 };
	WCHAR strpath[MAX_PATH] = { 0 };
	WCHAR strname[MAX_PATH] = { 0 };
	TCHAR* p = NULL;
	int nLen = 0;

	USES_CONVERSION;
	ZeroMemory(&file, sizeof(OPENFILENAME));
	file.lStructSize = sizeof(OPENFILENAME);
	file.lpstrFilter = _T("所有文件\0*.*\0\0");
	file.nFilterIndex = 1;
	file.lpstrFile = strfile;
	file.nMaxFile = sizeof(strfile);
	file.Flags = OFN_EXPLORER | OFN_ALLOWMULTISELECT;

	if (GetOpenFileName(&file)) {
		lstrcpyn(strpath, strfile, file.nFileOffset);
		strpath[file.nFileOffset] = '\0';
		nLen = lstrlen(strpath);

		if (strpath[nLen - 1] != '\\') {
			lstrcat(strpath, _T("\\"));
		}

		p = strfile + file.nFileOffset;
		while (*p) {
			ZeroMemory(strname, sizeof(strname));
			lstrcat(strname, strpath);
			lstrcat(strname, p);

			char chOriginFile[MAX_PATH] = { 0 };
			char chOriginName[MAX_PATH] = { 0 };
			char* pTemp = NULL;

			strcpy_s(chOriginFile, T2A(strname));
			pTemp = strrchr(chOriginFile, '\\');
			strcpy_s(chOriginName, ++pTemp);

			bool bRepeat = false;
			for (auto iter = m_vecComp.begin(); iter != m_vecComp.end(); ++iter) {
				if (!strcmp(iter->chPath, T2A(strname))) {
					bRepeat = true;
					break;
				}
			}

			if (!bRepeat) {
				TCompInfo sCompInfo = { 0 };
				sCompInfo.nSerial = m_vecComp.size() + 1;
				strcpy_s(sCompInfo.chName, chOriginName);
				strcpy_s(sCompInfo.chPath, T2A(strname));
				m_vecComp.push_back(sCompInfo);
			}
			p += lstrlen(p) + 1;
		}

		::PostMessageA(this->GetHWND(), WM_USER_MESSAGE_COMP_SEARCH, (WPARAM)0, (LPARAM)0);
	}
}

//----------------------------------------------
// @Function:	OnLButtonClickedCompDelBtn()
// @Purpose: CFrameMain click compress delete
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedCompDelBtn() {
	int nItem = m_pCompList->GetCurSel();
	if (nItem < 0) {
		MessageBoxW(this->GetHWND(), _T("请选中一条文件信息!"), _T("提示"), MB_OK | MB_ICONASTERISK);
		return;
	}

	m_vecComp.erase(m_vecComp.begin() + nItem);
	for (auto iter = m_vecComp.begin(); iter != m_vecComp.end(); ++iter) {
		iter->nSerial = iter - m_vecComp.begin() + 1;
	}

	::PostMessageA(this->GetHWND(), WM_USER_MESSAGE_COMP_SEARCH, (WPARAM)0, (LPARAM)0);
}

//----------------------------------------------
// @Function:	OnLButtonClickedCompClrBtn()
// @Purpose: CFrameMain click compress clear
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedCompClrBtn() {
	m_vecComp.clear();

	::PostMessageA(this->GetHWND(), WM_USER_MESSAGE_COMP_SEARCH, (WPARAM)0, (LPARAM)0);
}

//----------------------------------------------
// @Function:	OnLButtonClickedCompExportBtn()
// @Purpose: CFrameMain click compress export
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedCompExportBtn() {
	// check import files...
	if (m_vecComp.empty()) {
		MessageBoxA(this->GetHWND(), "请添加至少一个压缩文件!", "提示", MB_OK | MB_ICONASTERISK);
		return;
	}

	USES_CONVERSION;

	// check encrypt type...
	CDuiString strCompType = m_pCompTypeEdt->GetText();
	string strTypes[] = { "tar", "tar.gz", "zip" };
	vector<string> vecCompTypes(strTypes, strTypes + 3);
	bool bType = false;

	for (auto iter = vecCompTypes.begin(); iter != vecCompTypes.end(); ++iter) {
		if (!strcmp(iter->c_str(), T2A(strCompType.GetData()))) {
			bType = true;
			break;
		}
	}

	if (bType == false) {
		MessageBoxA(this->GetHWND(), "请填入支持的压缩类型(tar, tar.gz, zip)!", "提示", MB_OK | MB_ICONASTERISK);
		return;
	}

	// save file name...
	OPENFILENAME file;
	WCHAR strfile[MAX_PATH] = { 0 };

	wcscpy_s(strfile, A2T("undefine."));
	wcscat_s(strfile, strCompType.GetData());

	ZeroMemory(&file, sizeof(OPENFILENAME));
	file.lStructSize = sizeof(OPENFILENAME);
	file.lpstrFilter = _T("所有文件\0*.*\0\0");
	file.nFilterIndex = 1;
	file.lpstrFile = strfile;
	file.nMaxFile = sizeof(strfile);
	file.Flags = OFN_PATHMUSTEXIST | OFN_OVERWRITEPROMPT;

	if (GetSaveFileName(&file)) {
		m_pCompPathEdt->SetText(strfile);
	}
}

//----------------------------------------------
// @Function:	OnLButtonClickedCompStartBtn()
// @Purpose: CFrameMain click compress start
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedCompStartBtn() {
	// check import files...
	if (m_vecComp.empty()) {
		MessageBoxA(this->GetHWND(), "请添加至少一个压缩文件!", "提示", MB_OK | MB_ICONASTERISK);
		return;
	}

	USES_CONVERSION;

	// check encrypt type...
	CDuiString strCompType = m_pCompTypeEdt->GetText();
	string strTypes[] = { "tar", "tar.gz", "zip" };
	vector<string> vecCompTypes(strTypes, strTypes + 3);
	bool bType = false;

	for (auto iter = vecCompTypes.begin(); iter != vecCompTypes.end(); ++iter) {
		if (!strcmp(iter->c_str(), T2A(strCompType.GetData()))) {
			bType = true;
			break;
		}
	}

	if (bType == false) {
		MessageBoxA(this->GetHWND(), "请填入支持的压缩类型(tar, tar.gz, zip)!", "提示", MB_OK | MB_ICONASTERISK);
		return;
	}

	// check export compress...
	CDuiString strCompPath = m_pCompPathEdt->GetText();
	if (!strcmp("", T2A(strCompPath.GetData()))) {
		MessageBoxA(this->GetHWND(), "导出文件路径不能为空!", "警告", MB_OK | MB_ICONWARNING);
		return;
	}

	// set progress...
	m_pCompProgress->SetMinValue(0);
	m_pCompProgress->SetMaxValue(100);
	m_pCompProgress->SetValue(0);

	// reset timer...
	::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_COMP);
	//::SetTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_COMP, CONST_PROGRESS_REFRESH_TIME, NULL);

	// organize http request...
	CDuiString strCompJson;
	strCompJson = SpliceCompRequestJson(strCompType, strCompPath);

	// create thread handle request...
	HANDLE hThread = NULL;
	DWORD dwThreadID = 0;

	// create thread...
	hThread = ::CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)(&CFrameMain::OnPostCompRequestProcess), (LPVOID)(T2A(strCompJson.GetData())), 0, &dwThreadID);
	::CloseHandle(hThread);
}

//----------------------------------------------
// @Function:	OnLButtonClickedDecompDetialBtn()
// @Purpose: CFrameMain click decompress detial
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedDecompDetialBtn() {
}

//----------------------------------------------
// @Function:	OnLButtonClickedDecompImportBtn()
// @Purpose: CFrameMain click decompress import
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedDecompImportBtn() {
	OPENFILENAME file;
	WCHAR strfile[MAX_PATH] = { 0 };

	ZeroMemory(&file, sizeof(OPENFILENAME));
	file.lStructSize = sizeof(OPENFILENAME);
	file.lpstrFilter = _T("所有文件\0*.*\0\0");
	file.nFilterIndex = 1;
	file.lpstrFile = strfile;
	file.nMaxFile = sizeof(strfile);
	file.Flags = OFN_FILEMUSTEXIST | OFN_PATHMUSTEXIST;

	if (GetOpenFileName(&file)) {
		m_pDecompSrcEdt->SetText(strfile);
	}
}

//----------------------------------------------
// @Function:	OnLButtonClickedDecompExportBtn()
// @Purpose: CFrameMain click decompress export
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedDecompExportBtn() {
	CDuiString strImport = _T("");
	WCHAR strfile[MAX_PATH] = { 0 };

	USES_CONVERSION;

	wcscpy_s(strfile, A2T("undefine.tar.gz"));

	strImport = m_pDecompSrcEdt->GetText();
	if (!strcmp(T2A(strImport.GetData()), "")) {
		MessageBoxA(this->GetHWND(), "请选择源文件路径!", "提示", MB_OK | MB_ICONASTERISK);
		return;
	}

	OPENFILENAME file;

	ZeroMemory(&file, sizeof(OPENFILENAME));
	file.lStructSize = sizeof(OPENFILENAME);
	file.lpstrFilter = _T("所有文件\0*.*\0\0");
	file.nFilterIndex = 1;
	file.lpstrFile = strfile;
	file.nMaxFile = sizeof(strfile);
	file.Flags = OFN_PATHMUSTEXIST | OFN_OVERWRITEPROMPT;

	if (GetSaveFileName(&file)) {
		CDuiString dest(strfile);
		dest = dest.Left(dest.ReverseFind(L'\\') + 1);
		m_pDecompDestEdt->SetText(dest.GetData());
	}
}

//----------------------------------------------
// @Function:	OnLButtonClickedDecompStartBtn()
// @Purpose: CFrameMain click decompress start
// @Since: v1.00a
// @Para: None
// @Return: None
//----------------------------------------------
void CFrameMain::OnLButtonClickedDecompStartBtn() {
	CDuiString strDecompJson;
	CDuiString strDecompSrc;
	CDuiString strDecompDest;
	CDuiString strDecompType;

	USES_CONVERSION;

	strDecompSrc = m_pDecompSrcEdt->GetText();
	strDecompDest = m_pDecompDestEdt->GetText();

	// check source file...
	if (strDecompSrc.IsEmpty()) {
		MessageBoxA(this->GetHWND(), "压缩文件路径不能为空!", "警告", MB_OK | MB_ICONWARNING);
		return;
	}

	// check destination file...
	if (strDecompDest.IsEmpty()) {
		MessageBoxA(this->GetHWND(), "解压文件路径不能为空!", "警告", MB_OK | MB_ICONWARNING);
		return;
	}

	// check type...
	int pos = 0;
	pos = strDecompSrc.ReverseFind(L'.');
	if (pos == -1) {
		MessageBoxA(this->GetHWND(), "压缩文件解析错误!", "警告", MB_OK | MB_ICONWARNING);
		return;
	}

	strDecompType = strDecompSrc.Right(strDecompSrc.GetLength() - pos);
	if (!strDecompType.Compare(L".gz")) {
		strDecompType = L"tar.gz";
	} else if (!strDecompType.Compare(L".tar")) {
		strDecompType = L"tar";
	} else if (!strDecompType.Compare(L".zip")) {
		strDecompType = L"zip";
	} else {
		MessageBoxA(this->GetHWND(), "压缩文件类型错误!", "警告", MB_OK | MB_ICONWARNING);
		return;
	}

	// organize http request...
	strDecompJson = SpliceDecompRequestJson(strDecompSrc, strDecompDest, strDecompType);

	// set progress...
	m_pDecompProgress->SetMinValue(0);
	m_pDecompProgress->SetMaxValue(100);
	m_pDecompProgress->SetValue(0);

	// reset timer...
	::KillTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_DECOMP);
	//::SetTimer(this->GetHWND(), TIM_PROGRESS_REFRESH_DECOMP, CONST_PROGRESS_REFRESH_TIME, NULL);

	// create thread handle request...
	HANDLE hThread = NULL;
	DWORD dwThreadID = 0;

	// create thread...
	hThread = ::CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)(&CFrameMain::OnPostDecompRequestProcess), (LPVOID)(T2A(strDecompJson.GetData())), 0, &dwThreadID);
	::CloseHandle(hThread);
}
