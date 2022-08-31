// Code generated by "stringer -output func_string.go -type=BuiltinFunc"; DO NOT EDIT.

package asm

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FnUnspec-0]
	_ = x[FnMapLookupElem-1]
	_ = x[FnMapUpdateElem-2]
	_ = x[FnMapDeleteElem-3]
	_ = x[FnProbeRead-4]
	_ = x[FnKtimeGetNs-5]
	_ = x[FnTracePrintk-6]
	_ = x[FnGetPrandomU32-7]
	_ = x[FnGetSmpProcessorId-8]
	_ = x[FnSkbStoreBytes-9]
	_ = x[FnL3CsumReplace-10]
	_ = x[FnL4CsumReplace-11]
	_ = x[FnTailCall-12]
	_ = x[FnCloneRedirect-13]
	_ = x[FnGetCurrentPidTgid-14]
	_ = x[FnGetCurrentUidGid-15]
	_ = x[FnGetCurrentComm-16]
	_ = x[FnGetCgroupClassid-17]
	_ = x[FnSkbVlanPush-18]
	_ = x[FnSkbVlanPop-19]
	_ = x[FnSkbGetTunnelKey-20]
	_ = x[FnSkbSetTunnelKey-21]
	_ = x[FnPerfEventRead-22]
	_ = x[FnRedirect-23]
	_ = x[FnGetRouteRealm-24]
	_ = x[FnPerfEventOutput-25]
	_ = x[FnSkbLoadBytes-26]
	_ = x[FnGetStackid-27]
	_ = x[FnCsumDiff-28]
	_ = x[FnSkbGetTunnelOpt-29]
	_ = x[FnSkbSetTunnelOpt-30]
	_ = x[FnSkbChangeProto-31]
	_ = x[FnSkbChangeType-32]
	_ = x[FnSkbUnderCgroup-33]
	_ = x[FnGetHashRecalc-34]
	_ = x[FnGetCurrentTask-35]
	_ = x[FnProbeWriteUser-36]
	_ = x[FnCurrentTaskUnderCgroup-37]
	_ = x[FnSkbChangeTail-38]
	_ = x[FnSkbPullData-39]
	_ = x[FnCsumUpdate-40]
	_ = x[FnSetHashInvalid-41]
	_ = x[FnGetNumaNodeId-42]
	_ = x[FnSkbChangeHead-43]
	_ = x[FnXdpAdjustHead-44]
	_ = x[FnProbeReadStr-45]
	_ = x[FnGetSocketCookie-46]
	_ = x[FnGetSocketUid-47]
	_ = x[FnSetHash-48]
	_ = x[FnSetsockopt-49]
	_ = x[FnSkbAdjustRoom-50]
	_ = x[FnRedirectMap-51]
	_ = x[FnSkRedirectMap-52]
	_ = x[FnSockMapUpdate-53]
	_ = x[FnXdpAdjustMeta-54]
	_ = x[FnPerfEventReadValue-55]
	_ = x[FnPerfProgReadValue-56]
	_ = x[FnGetsockopt-57]
	_ = x[FnOverrideReturn-58]
	_ = x[FnSockOpsCbFlagsSet-59]
	_ = x[FnMsgRedirectMap-60]
	_ = x[FnMsgApplyBytes-61]
	_ = x[FnMsgCorkBytes-62]
	_ = x[FnMsgPullData-63]
	_ = x[FnBind-64]
	_ = x[FnXdpAdjustTail-65]
	_ = x[FnSkbGetXfrmState-66]
	_ = x[FnGetStack-67]
	_ = x[FnSkbLoadBytesRelative-68]
	_ = x[FnFibLookup-69]
	_ = x[FnSockHashUpdate-70]
	_ = x[FnMsgRedirectHash-71]
	_ = x[FnSkRedirectHash-72]
	_ = x[FnLwtPushEncap-73]
	_ = x[FnLwtSeg6StoreBytes-74]
	_ = x[FnLwtSeg6AdjustSrh-75]
	_ = x[FnLwtSeg6Action-76]
	_ = x[FnRcRepeat-77]
	_ = x[FnRcKeydown-78]
	_ = x[FnSkbCgroupId-79]
	_ = x[FnGetCurrentCgroupId-80]
	_ = x[FnGetLocalStorage-81]
	_ = x[FnSkSelectReuseport-82]
	_ = x[FnSkbAncestorCgroupId-83]
	_ = x[FnSkLookupTcp-84]
	_ = x[FnSkLookupUdp-85]
	_ = x[FnSkRelease-86]
	_ = x[FnMapPushElem-87]
	_ = x[FnMapPopElem-88]
	_ = x[FnMapPeekElem-89]
	_ = x[FnMsgPushData-90]
	_ = x[FnMsgPopData-91]
	_ = x[FnRcPointerRel-92]
	_ = x[FnSpinLock-93]
	_ = x[FnSpinUnlock-94]
	_ = x[FnSkFullsock-95]
	_ = x[FnTcpSock-96]
	_ = x[FnSkbEcnSetCe-97]
	_ = x[FnGetListenerSock-98]
	_ = x[FnSkcLookupTcp-99]
	_ = x[FnTcpCheckSyncookie-100]
	_ = x[FnSysctlGetName-101]
	_ = x[FnSysctlGetCurrentValue-102]
	_ = x[FnSysctlGetNewValue-103]
	_ = x[FnSysctlSetNewValue-104]
	_ = x[FnStrtol-105]
	_ = x[FnStrtoul-106]
	_ = x[FnSkStorageGet-107]
	_ = x[FnSkStorageDelete-108]
	_ = x[FnSendSignal-109]
	_ = x[FnTcpGenSyncookie-110]
}

const _BuiltinFunc_name = "FnUnspecFnMapLookupElemFnMapUpdateElemFnMapDeleteElemFnProbeReadFnKtimeGetNsFnTracePrintkFnGetPrandomU32FnGetSmpProcessorIdFnSkbStoreBytesFnL3CsumReplaceFnL4CsumReplaceFnTailCallFnCloneRedirectFnGetCurrentPidTgidFnGetCurrentUidGidFnGetCurrentCommFnGetCgroupClassidFnSkbVlanPushFnSkbVlanPopFnSkbGetTunnelKeyFnSkbSetTunnelKeyFnPerfEventReadFnRedirectFnGetRouteRealmFnPerfEventOutputFnSkbLoadBytesFnGetStackidFnCsumDiffFnSkbGetTunnelOptFnSkbSetTunnelOptFnSkbChangeProtoFnSkbChangeTypeFnSkbUnderCgroupFnGetHashRecalcFnGetCurrentTaskFnProbeWriteUserFnCurrentTaskUnderCgroupFnSkbChangeTailFnSkbPullDataFnCsumUpdateFnSetHashInvalidFnGetNumaNodeIdFnSkbChangeHeadFnXdpAdjustHeadFnProbeReadStrFnGetSocketCookieFnGetSocketUidFnSetHashFnSetsockoptFnSkbAdjustRoomFnRedirectMapFnSkRedirectMapFnSockMapUpdateFnXdpAdjustMetaFnPerfEventReadValueFnPerfProgReadValueFnGetsockoptFnOverrideReturnFnSockOpsCbFlagsSetFnMsgRedirectMapFnMsgApplyBytesFnMsgCorkBytesFnMsgPullDataFnBindFnXdpAdjustTailFnSkbGetXfrmStateFnGetStackFnSkbLoadBytesRelativeFnFibLookupFnSockHashUpdateFnMsgRedirectHashFnSkRedirectHashFnLwtPushEncapFnLwtSeg6StoreBytesFnLwtSeg6AdjustSrhFnLwtSeg6ActionFnRcRepeatFnRcKeydownFnSkbCgroupIdFnGetCurrentCgroupIdFnGetLocalStorageFnSkSelectReuseportFnSkbAncestorCgroupIdFnSkLookupTcpFnSkLookupUdpFnSkReleaseFnMapPushElemFnMapPopElemFnMapPeekElemFnMsgPushDataFnMsgPopDataFnRcPointerRelFnSpinLockFnSpinUnlockFnSkFullsockFnTcpSockFnSkbEcnSetCeFnGetListenerSockFnSkcLookupTcpFnTcpCheckSyncookieFnSysctlGetNameFnSysctlGetCurrentValueFnSysctlGetNewValueFnSysctlSetNewValueFnStrtolFnStrtoulFnSkStorageGetFnSkStorageDeleteFnSendSignalFnTcpGenSyncookie"

var _BuiltinFunc_index = [...]uint16{0, 8, 23, 38, 53, 64, 76, 89, 104, 123, 138, 153, 168, 178, 193, 212, 230, 246, 264, 277, 289, 306, 323, 338, 348, 363, 380, 394, 406, 416, 433, 450, 466, 481, 497, 512, 528, 544, 568, 583, 596, 608, 624, 639, 654, 669, 683, 700, 714, 723, 735, 750, 763, 778, 793, 808, 828, 847, 859, 875, 894, 910, 925, 939, 952, 958, 973, 990, 1000, 1022, 1033, 1049, 1066, 1082, 1096, 1115, 1133, 1148, 1158, 1169, 1182, 1202, 1219, 1238, 1259, 1272, 1285, 1296, 1309, 1321, 1334, 1347, 1359, 1373, 1383, 1395, 1407, 1416, 1429, 1446, 1460, 1479, 1494, 1517, 1536, 1555, 1563, 1572, 1586, 1603, 1615, 1632}

func (i BuiltinFunc) String() string {
	if i < 0 || i >= BuiltinFunc(len(_BuiltinFunc_index)-1) {
		return "BuiltinFunc(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BuiltinFunc_name[_BuiltinFunc_index[i]:_BuiltinFunc_index[i+1]]
}
