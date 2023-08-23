// AUTO-GENERATED CODE, DO NOT MODIFY

package coreaudiotypes

// An integer type for audio operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/avaudiointeger?language=objc
type AudioInteger int32

// Codes that describe error conditions that may occur when performing audio session operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/avaudiosessionerrorcode?language=objc
type AudioSessionErrorCode AudioInteger

const (
	AudioSessionErrorCodeBadParam              AudioSessionErrorCode = -50
	AudioSessionErrorCodeCannotInterruptOthers AudioSessionErrorCode = 560557684
	AudioSessionErrorCodeCannotStartPlaying    AudioSessionErrorCode = 561015905
	AudioSessionErrorCodeCannotStartRecording  AudioSessionErrorCode = 561145187
	AudioSessionErrorCodeExpiredSession        AudioSessionErrorCode = 561210739
	AudioSessionErrorCodeIncompatibleCategory  AudioSessionErrorCode = 560161140
	AudioSessionErrorCodeIsBusy                AudioSessionErrorCode = 560030580
	AudioSessionErrorCodeMediaServicesFailed   AudioSessionErrorCode = 1836282486
	AudioSessionErrorCodeMissingEntitlement    AudioSessionErrorCode = 1701737535
	AudioSessionErrorCodeNone                  AudioSessionErrorCode = 0
	AudioSessionErrorCodeResourceNotAvailable  AudioSessionErrorCode = 561145203
	AudioSessionErrorCodeSessionNotActive      AudioSessionErrorCode = 1768841571
	AudioSessionErrorCodeSiriIsRecording       AudioSessionErrorCode = 1936290409
	AudioSessionErrorCodeUnspecified           AudioSessionErrorCode = 2003329396
)

// An unsigned integer type for audio operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/avaudiouinteger?language=objc
type AudioUInteger int32

// The supported channel bitmaps to use when defining channel layouts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiochannelbitmap?language=objc
type ChannelBitmap uint32

const (
	KChannelBit_Center               ChannelBitmap = 4
	KChannelBit_CenterSurround       ChannelBitmap = 256
	KChannelBit_CenterTopFront       ChannelBitmap = 8192
	KChannelBit_CenterTopMiddle      ChannelBitmap = 2048
	KChannelBit_CenterTopRear        ChannelBitmap = 33554432
	KChannelBit_LFEScreen            ChannelBitmap = 8
	KChannelBit_Left                 ChannelBitmap = 1
	KChannelBit_LeftCenter           ChannelBitmap = 64
	KChannelBit_LeftSurround         ChannelBitmap = 16
	KChannelBit_LeftSurroundDirect   ChannelBitmap = 512
	KChannelBit_LeftTopFront         ChannelBitmap = 4096
	KChannelBit_LeftTopMiddle        ChannelBitmap = 2097152
	KChannelBit_LeftTopRear          ChannelBitmap = 16777216
	KChannelBit_Right                ChannelBitmap = 2
	KChannelBit_RightCenter          ChannelBitmap = 128
	KChannelBit_RightSurround        ChannelBitmap = 32
	KChannelBit_RightSurroundDirect  ChannelBitmap = 1024
	KChannelBit_RightTopFront        ChannelBitmap = 16384
	KChannelBit_RightTopMiddle       ChannelBitmap = 8388608
	KChannelBit_RightTopRear         ChannelBitmap = 67108864
	KChannelBit_TopBackCenter        ChannelBitmap = 65536
	KChannelBit_TopBackLeft          ChannelBitmap = 32768
	KChannelBit_TopBackRight         ChannelBitmap = 131072
	KChannelBit_TopCenterSurround    ChannelBitmap = 2048
	KChannelBit_VerticalHeightCenter ChannelBitmap = 8192
	KChannelBit_VerticalHeightLeft   ChannelBitmap = 4096
	KChannelBit_VerticalHeightRight  ChannelBitmap = 16384
)

// Indexes the fields of the mCoordinates array in an AudioChannelDescription structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiochannelcoordinateindex?language=objc
type ChannelCoordinateIndex uint32

const (
	KChannelCoordinates_Azimuth   ChannelCoordinateIndex = 0
	KChannelCoordinates_BackFront ChannelCoordinateIndex = 1
	KChannelCoordinates_Distance  ChannelCoordinateIndex = 2
	KChannelCoordinates_DownUp    ChannelCoordinateIndex = 2
	KChannelCoordinates_Elevation ChannelCoordinateIndex = 1
	KChannelCoordinates_LeftRight ChannelCoordinateIndex = 0
)

// Constants that define the audio channel flags of an audio channel description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiochannelflags?language=objc
type ChannelFlags uint32

const (
	KChannelFlags_AllOff                 ChannelFlags = 0
	KChannelFlags_Meters                 ChannelFlags = 4
	KChannelFlags_RectangularCoordinates ChannelFlags = 1
	KChannelFlags_SphericalCoordinates   ChannelFlags = 2
)

// Identifies how an audio data channel is to be used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiochannellabel?language=objc
type ChannelLabel uint32

// Identifies a previously-defined channel layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiochannellayouttag?language=objc
type ChannelLayoutTag uint32

// A type definition for audio format flags. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audioformatflags?language=objc
type FormatFlags uint32

// A type definition for audio format identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audioformatid?language=objc
type FormatID uint32

// The canonical audio data sample type for input and output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiosampletype?language=objc
type SampleType int16

// A unique identifier of an audio session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiosessionid?language=objc
type SessionID uint32

// A structure that represents flags for a timestamp. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiotimestampflags?language=objc
type TimeStampFlags uint32

const (
	KTimeStampHostTimeValid       TimeStampFlags = 2
	KTimeStampNothingValid        TimeStampFlags = 0
	KTimeStampRateScalarValid     TimeStampFlags = 4
	KTimeStampSMPTETimeValid      TimeStampFlags = 16
	KTimeStampSampleHostTimeValid TimeStampFlags = 3
	KTimeStampSampleTimeValid     TimeStampFlags = 1
	KTimeStampWordClockTimeValid  TimeStampFlags = 8
)

// The canonical audio data sample type for audio processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/audiounitsampletype?language=objc
type UnitSampleType int32

// Constants that define the type of MPEG-4 audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/mpeg4objectid?language=objc
type MPEG4ObjectID int32

const (
	kMPEG4Object_AAC_LC       MPEG4ObjectID = 2
	kMPEG4Object_AAC_LTP      MPEG4ObjectID = 4
	kMPEG4Object_AAC_Main     MPEG4ObjectID = 1
	kMPEG4Object_AAC_SBR      MPEG4ObjectID = 5
	kMPEG4Object_AAC_SSR      MPEG4ObjectID = 3
	kMPEG4Object_AAC_Scalable MPEG4ObjectID = 6
	kMPEG4Object_CELP         MPEG4ObjectID = 8
	kMPEG4Object_HVXC         MPEG4ObjectID = 9
	kMPEG4Object_TwinVQ       MPEG4ObjectID = 7
)

// A structure that defines SMPTE time flags. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/smptetimeflags?language=objc
type SMPTETimeFlags uint32

const (
	kSMPTETimeRunning SMPTETimeFlags = 2
	kSMPTETimeUnknown SMPTETimeFlags = 0
	kSMPTETimeValid   SMPTETimeFlags = 1
)

// Constants that define SMPTE time types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/smptetimetype?language=objc
type SMPTETimeType uint32

const (
	kSMPTETimeType2398     SMPTETimeType = 11
	kSMPTETimeType24       SMPTETimeType = 0
	kSMPTETimeType25       SMPTETimeType = 1
	kSMPTETimeType2997     SMPTETimeType = 4
	kSMPTETimeType2997Drop SMPTETimeType = 5
	kSMPTETimeType30       SMPTETimeType = 3
	kSMPTETimeType30Drop   SMPTETimeType = 2
	kSMPTETimeType50       SMPTETimeType = 10
	kSMPTETimeType5994     SMPTETimeType = 7
	kSMPTETimeType5994Drop SMPTETimeType = 9
	kSMPTETimeType60       SMPTETimeType = 6
	kSMPTETimeType60Drop   SMPTETimeType = 8
)
