/*
*     COPYRIGHT NOTICE
*     Copyright(c) 2017~2019, Team Gorgeous Bubble
*     All rights reserved.
*
* @file		frame_main.h
* @brief	The Detector Project
* @author	Alopex/Aurora
* @version	v1.00a
* @date		2019-12-29
*/
#pragma once

#ifndef __FRAME_MAIN_H_
#define __FRAME_MAIN_H_

// Include DUI frame header files
#include "common_wnd.h"
#include "types.h"

// Class definition
class CFrameMain : public CWindowWnd, public INotifyUI {
public:
	virtual LPCTSTR GetWindowClassName() const;
	virtual UINT GetClassStyle() const;
	virtual void Notify(TNotifyUI& msg);
	virtual LRESULT HandleMessage(UINT uMsg, WPARAM wParam, LPARAM lParam);

protected:
	CPaintManagerUI m_PaintManager;

public:
	LRESULT OnCreate(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnLButtonDown(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnClose(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnTimer(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnNcActivate(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnNcCalcSize(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnNcPaint(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnNcHitTest(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnSize(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnGetMinMaxInfo(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnSysCommand(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);

public:
	LRESULT OnUserMessagePacketSearch(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnUserMessagePacketAddItem(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);

public:
	// Main Layout...
	CButtonUI* m_pCloseBtn;
	CButtonUI* m_pRestoreBtn;
	CButtonUI* m_pMaxBtn;
	CButtonUI* m_pMinBtn;

	// Tab Layout...
	CTabLayoutUI* m_pMainTab;
	COptionUI* m_pPackOpt;
	COptionUI* m_pUnpackOpt;
	COptionUI* m_pCompOpt;
	COptionUI* m_pDecompOpt;	
	COptionUI* m_pBaseOpt;
	COptionUI* m_pAppOpt;
	COptionUI* m_pMoreOpt;
	COptionUI* m_pAboutOpt;

	// Packet...
	CButtonUI* m_pPackAddBtn;
	CButtonUI* m_pPackDelBtn;
	CListUI* m_pPackList;
	CEditUI* m_pPackTypeEdt;
	CEditUI* m_pPackPathEdt;
	CButtonUI* m_pPackExportBtn;
	CProgressUI* m_pPackProgress;
	CButtonUI* m_pPackStartBtn;

public:
	CPaintManagerUI& GetPaintManager();

private:
	HMENU m_hMenu;
	NOTIFYICONDATA m_nid;

public:
	vector<TPacketInfo> m_vecPacket;

public:
	void ConstructExtra();
	void DestructExtra();
	void InitMenuShow();
	void InitWindowSharp();
	void InitControls();

	static DWORD CALLBACK OnSearchPacketItemsProcess(LPVOID lpParameter);

public:
	LRESULT OnUserMessageMenu(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);

public:
	void OnLButtonClickedMinBtn();
	void OnLButtonClickedMaxBtn();
	void OnLButtonClickedRestoreBtn();
	void OnLButtonClickedCloseBtn();

	void OnLButtonClickedPacketAddBtn();
	void OnLButtonClickedPacketDelBtn();
	void OnLButtonClickedPacketExportBtn();
	void OnLButtonClickedPacketStartBtn();
};

// External
extern CFrameMain* g_pFrameMain;

#endif // !__FRAME_MAIN_H_

