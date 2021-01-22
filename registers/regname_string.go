// Code generated by "stringer -type=RegName"; DO NOT EDIT.

package registers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ModelNumber-0]
	_ = x[FirmwareVersion-1]
	_ = x[ServoID-2]
	_ = x[BaudRate-3]
	_ = x[ReturnDelayTime-4]
	_ = x[CwAngleLimit-5]
	_ = x[CcwAngleLimit-6]
	_ = x[HighestLimitTemperature-7]
	_ = x[LowestLimitVoltage-8]
	_ = x[HighestLimitVoltage-9]
	_ = x[MaxTorque-10]
	_ = x[StatusReturnLevel-11]
	_ = x[AlarmLed-12]
	_ = x[AlarmShutdown-13]
	_ = x[TorqueEnable-14]
	_ = x[Led-15]
	_ = x[CwComplianceMargin-16]
	_ = x[CcwComplianceMargin-17]
	_ = x[CwComplianceSlope-18]
	_ = x[CcwComplianceSlope-19]
	_ = x[GoalPosition-20]
	_ = x[MovingSpeed-21]
	_ = x[TorqueLimit-22]
	_ = x[PresentPosition-23]
	_ = x[PresentSpeed-24]
	_ = x[PresentLoad-25]
	_ = x[PresentVoltage-26]
	_ = x[PresentTemperature-27]
	_ = x[RegisteredInstruction-28]
	_ = x[Moving-29]
	_ = x[Lock-30]
	_ = x[Punch-31]
	_ = x[ControlMode-32]
	_ = x[DGain-33]
	_ = x[IGain-34]
	_ = x[PGain-35]
	_ = x[HardwareErrorStatus-36]
	_ = x[ModelInfo-37]
	_ = x[DriveMode-38]
	_ = x[OperatingMode-39]
	_ = x[SecondaryID-40]
	_ = x[ProtocolType-41]
	_ = x[HomingOffset-42]
	_ = x[MovingThreshold-43]
	_ = x[MaxLimitVoltage-44]
	_ = x[MinLimitVoltage-45]
	_ = x[PWMLimit-46]
	_ = x[VelocityLimit-47]
	_ = x[MaxPositionLimit-48]
	_ = x[MinPositionLimit-49]
	_ = x[VIGain-50]
	_ = x[VDGain-51]
	_ = x[Feedforward2ndGain-52]
	_ = x[Feedforward1stGain-53]
	_ = x[BusWatchdog-54]
	_ = x[GoalPWM-55]
	_ = x[ProfileAcceleration-56]
	_ = x[ProfileVelocity-57]
	_ = x[RealtimeTick-58]
	_ = x[MovingStatus-59]
	_ = x[PresentPWM-60]
	_ = x[PresentVelocity-61]
	_ = x[VelocityTrajectory-62]
	_ = x[PositionTrajectory-63]
	_ = x[CurrentLimit-64]
	_ = x[GoalCurrent-65]
	_ = x[PresentCurrent-66]
	_ = x[GoalVelocity-67]
	_ = x[GoalTorque-68]
}

const _RegName_name = "ModelNumberFirmwareVersionServoIDBaudRateReturnDelayTimeCwAngleLimitCcwAngleLimitHighestLimitTemperatureLowestLimitVoltageHighestLimitVoltageMaxTorqueStatusReturnLevelAlarmLedAlarmShutdownTorqueEnableLedCwComplianceMarginCcwComplianceMarginCwComplianceSlopeCcwComplianceSlopeGoalPositionMovingSpeedTorqueLimitPresentPositionPresentSpeedPresentLoadPresentVoltagePresentTemperatureRegisteredInstructionMovingLockPunchControlModeDGainIGainPGainHardwareErrorStatusModelInfoDriveModeOperatingModeSecondaryIDProtocolTypeHomingOffsetMovingThresholdMaxLimitVoltageMinLimitVoltagePWMLimitVelocityLimitMaxPositionLimitMinPositionLimitVIGainVDGainFeedforward2ndGainFeedforward1stGainBusWatchdogGoalPWMProfileAccelerationProfileVelocityRealtimeTickMovingStatusPresentPWMPresentVelocityVelocityTrajectoryPositionTrajectoryCurrentLimitGoalCurrentPresentCurrentGoalVelocityGoalTorque"

var _RegName_index = [...]uint16{0, 11, 26, 33, 41, 56, 68, 81, 104, 122, 141, 150, 167, 175, 188, 200, 203, 221, 240, 257, 275, 287, 298, 309, 324, 336, 347, 361, 379, 400, 406, 410, 415, 426, 431, 436, 441, 460, 469, 478, 491, 502, 514, 526, 541, 556, 571, 579, 592, 608, 624, 630, 636, 654, 672, 683, 690, 709, 724, 736, 748, 758, 773, 791, 809, 821, 832, 846, 858, 868}

func (i RegName) String() string {
	if i < 0 || i >= RegName(len(_RegName_index)-1) {
		return "RegName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RegName_name[_RegName_index[i]:_RegName_index[i+1]]
}
