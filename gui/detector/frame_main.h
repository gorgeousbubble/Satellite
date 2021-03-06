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

// Include C/C++ header files
#include <iostream>

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
	LRESULT OnUserMessagePacketResult(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnUserMessageUnpackSearch(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnUserMessageUnpackAddItem(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnUserMessageUnpackResult(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnUserMessageCompSearch(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnUserMessageCompAddItem(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnUserMessageCompResult(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnUserMessageDecompSearch(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnUserMessageDecompAddItem(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);
	LRESULT OnUserMessageDecompResult(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);

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
	CButtonUI* m_pPackClrBtn;
	CListUI* m_pPackList;
	CEditUI* m_pPackTypeEdt;
	CEditUI* m_pPackPathEdt;
	CButtonUI* m_pPackExportBtn;
	CProgressUI* m_pPackProgress;
	CButtonUI* m_pPackStartBtn;

	// Unpack...
	CButtonUI* m_pUnpackSelBtn;
	CButtonUI* m_pUnpackUnselBtn;
	CButtonUI* m_pUnpackRetBtn;
	CButtonUI* m_pUnpackUnretBtn;
	CButtonUI* m_pUnpackDetBtn;
	CListUI* m_pUnpackList;
	CEditUI* m_pUnpackSrcEdt;
	CEditUI* m_pUnpackDestEdt;
	CButtonUI* m_pUnpackImportBtn;
	CButtonUI* m_pUnpackExportBtn;
	CProgressUI* m_pUnpackProgress;
	CButtonUI* m_pUnpackStartBtn;

	// Compress...
	CButtonUI* m_pCompAddBtn;
	CButtonUI* m_pCompDelBtn;
	CButtonUI* m_pCompClrBtn;
	CListUI* m_pCompList;
	CEditUI* m_pCompTypeEdt;
	CEditUI* m_pCompPathEdt;
	CButtonUI* m_pCompExportBtn;
	CProgressUI* m_pCompProgress;
	CButtonUI* m_pCompStartBtn;

	// Decompress...
	CButtonUI* m_pDecompDetBtn;
	CListUI* m_pDecompList;
	CEditUI* m_pDecompSrcEdt;
	CEditUI* m_pDecompDestEdt;
	CButtonUI* m_pDecompImportBtn;
	CButtonUI* m_pDecompExportBtn;
	CProgressUI* m_pDecompProgress;
	CButtonUI* m_pDecompStartBtn;

public:
	CPaintManagerUI& GetPaintManager();

private:
	HMENU m_hMenu;
	NOTIFYICONDATA m_nid;

public:
	vector<TPacketInfo> m_vecPacket;
	vector<TUnpackInfo> m_vecUnpack;
	vector<TCompInfo> m_vecComp;

public:
	void ConstructExtra();
	void DestructExtra();
	void InitMenuShow();
	void InitWindowSharp();
	void InitControls();
	void InitCurl();

	CDuiString SplicePackRequestJson(CDuiString strPacketType, CDuiString strPacketPath);
	CDuiString SplicePackProcessRequestJson(CDuiString strPacketType);
	CDuiString SpliceUnpackRequestJson(CDuiString strUnpackSrc, CDuiString strUnpackDest);
	CDuiString SpliceUnpackVerboseRequestJson(CDuiString strUnpackSrc);
	CDuiString SpliceUnpackProcessRequestJson(CDuiString strUnpackSrc);
	CDuiString SpliceUnpackOneFileRequestJson(CDuiString strUnpackSrc, CDuiString strUnpackDest, CDuiString strUnpackFile);
	CDuiString SpliceCompRequestJson(CDuiString strCompType, CDuiString strCompPath);
	CDuiString SpliceDecompRequestJson(CDuiString strDecompSrc, CDuiString strDecompDest, CDuiString strDecompType);

	int GetValueFromResponse(string result, string label_first, string label_last, string& value);
	int GetValueFromResponseWithPos(string result, string label_first, string label_last, string& value, int& poi);

	static DWORD CALLBACK OnSearchPacketItemsProcess(LPVOID lpParameter);
	static DWORD CALLBACK OnSearchUnpackItemsProcess(LPVOID lpParameter);
	static DWORD CALLBACK OnSearchCompItemsProcess(LPVOID lpParameter);
	static DWORD CALLBACK OnPostPackRequestProcess(LPVOID lpParameter);
	static DWORD CALLBACK OnPostUnpackRequestProcess(LPVOID lpParameter);
	static DWORD CALLBACK OnPostUnpackConfineRequestProcess(LPVOID lpParameter);
	static DWORD CALLBACK OnPostUnpackFileRequestProcess(LPVOID lpParameter);
	static DWORD CALLBACK OnPostUnpackFileConfineRequestProcess(LPVOID lpParameter);
	static DWORD CALLBACK OnPostCompRequestProcess(LPVOID lpParameter);
	static DWORD CALLBACK OnPostDecompRequestProcess(LPVOID lpParameter);

private:
	DWORD StringToDword(string value);
	void StringSplit(const std::string& s, std::vector<std::string>& v, const std::string& c);

public:
	LRESULT OnUserMessageMenu(UINT uMsg, WPARAM wParam, LPARAM lParam, BOOL& bHandled);

public:
	void OnLButtonClickedMinBtn();
	void OnLButtonClickedMaxBtn();
	void OnLButtonClickedRestoreBtn();
	void OnLButtonClickedCloseBtn();

	void OnLButtonClickedPacketAddBtn();
	void OnLButtonClickedPacketDelBtn();
	void OnLButtonClickedPacketClrBtn();
	void OnLButtonClickedPacketExportBtn();
	void OnLButtonClickedPacketStartBtn();

	void OnLButtonClickedUnpackSelectBtn();
	void OnLButtonClickedUnpackUnselectBtn();
	void OnLButtonClickedUnpackRestrictBtn();
	void OnLButtonClickedUnpackUnrestrictBtn();
	void OnLButtonClickedUnpackDetialBtn();
	void OnLButtonClickedUnpackImportBtn();
	void OnLButtonClickedUnpackExportBtn();
	void OnLButtonClickedUnpackStartBtn();

	void OnLButtonClickedCompAddBtn();
	void OnLButtonClickedCompDelBtn();
	void OnLButtonClickedCompClrBtn();
	void OnLButtonClickedCompExportBtn();
	void OnLButtonClickedCompStartBtn();

	void OnLButtonClickedDecompDetialBtn();
	void OnLButtonClickedDecompImportBtn();
	void OnLButtonClickedDecompExportBtn();
	void OnLButtonClickedDecompStartBtn();
};

// External
extern CFrameMain* g_pFrameMain;

#endif // !__FRAME_MAIN_H_
