package consts

import "github.com/pkg/errors"

var (
	// ErrSettingsNil gets returned when no settings are supplied to ComposeAPI().
	ErrSettingsNil = errors.New("settings must not be nil")
	// ErrInvalidSettingsType gets returned if the wrong underlying type of Settings were supplied for creating a Provider.
	ErrInvalidSettingsType = errors.New("incompatible settings type supplied")
	// ErrInconsistentSubtangle gets returned when the tail transaction is not consistent during promotion.
	ErrInconsistentSubtangle = errors.New("inconsistent subtangle")
	// ErrInvalidSqueezeLength gets returned when the squeeze length is not a multiple of 243.
	ErrInvalidSqueezeLength = errors.New("squeeze length must be a multiple of 243")
	// ErrInvalidBatchSize is returned when the batch size is invalid.
	ErrInvalidBatchSize = errors.New("invalid batch size")
	// ErrInvalidTritsLength gets returned when the trits length are invalid for the given operation.
	ErrInvalidTritsLength = errors.New("invalid trits length")
	// ErrInvalidTrytesLength gets returned when the trytes length are invalid for the given operation.
	ErrInvalidTrytesLength = errors.New("invalid trytes length")
	// ErrInvalidBytesLength gets returned when the bytes length are invalid for the given operation.
	ErrInvalidBytesLength = errors.New("invalid bytes length")
	// ErrInsufficientBalance gets returned when an operation needs a certain amount of balance to fulfill the operation.
	ErrInsufficientBalance = errors.New("insufficient balance")
	// ErrInvalidAddress gets returned for invalid address parameters.
	ErrInvalidAddress = errors.New("invalid address")
	// ErrInvalidRemainderAddress gets returned for invalid remainder address parameters.
	ErrInvalidRemainderAddress = errors.New("invalid remainder address")
	// ErrInvalidBranchTransaction gets returned for invalid branch transaction parameters.
	ErrInvalidBranchTransaction = errors.New("invalid branch transaction")
	// ErrInvalidBundle gets returned for Bundles which are schematically wrong or/and don't pass validation.
	ErrInvalidBundle = errors.New("invalid bundle")
	// ErrInvalidBundleTotalValue gets returned for Bundles whose total value would overflow/underflow from the total supply.
	ErrInvalidBundleTotalValue = errors.New("invalid bundle total value")
	// ErrInvalidBundleAddressValue gets returned for Bundles with balance mutations in an address that would overflow/underflow from the total supply.
	ErrInvalidBundleAddressValue = errors.New("invalid bundle address value")
	// ErrInvalidBundleHash gets returned for invalid bundle hash parameters.
	ErrInvalidBundleHash = errors.New("invalid bundle hash")
	// ErrInvalidSignature gets returned for bundles with invalid signatures.
	ErrInvalidSignature = errors.New("invalid signature")
	// ErrInvalidChecksum gets returned for addresses with invalid checksum.
	ErrInvalidChecksum = errors.New("invalid checksum")
	// ErrInvalidHash gets returned for invalid hash parameters.
	ErrInvalidHash = errors.New("invalid hash")
	// ErrInvalidValue gets returned for invalid values (e.g. upper 48 trits of the value field used).
	ErrInvalidValue = errors.New("invalid value")
	// ErrInvalidIndex gets returned for invalid index parameters.
	ErrInvalidIndex = errors.New("invalid index option")
	// ErrInvalidTotalOption gets returned for invalid total option parameters.
	ErrInvalidTotalOption = errors.New("invalid total option")
	// ErrInvalidInput gets returned for invalid input parameters.
	ErrInvalidInput = errors.New("invalid input")
	// ErrInvalidSecurityLevel gets returned for invalid security level parameters.
	ErrInvalidSecurityLevel = errors.New("invalid security option")
	// ErrInvalidSeed gets returned for invalid seed parameters.
	ErrInvalidSeed = errors.New("invalid seed")
	// ErrInvalidStartEndOptions gets returned for invalid end options.
	ErrInvalidStartEndOptions = errors.New("invalid end option")
	// ErrInvalidTag gets returned for invalid tags.
	ErrInvalidTag = errors.New("invalid tag")
	// ErrInvalidTransaction gets returned when transactions trits don't make up a valid transaction.
	ErrInvalidTransaction = errors.New("invalid transaction")
	// ErrInvalidTransactionTrytes gets returned for invalid transaction trytes.
	ErrInvalidTransactionTrytes = errors.New("invalid transaction trytes")
	// ErrInvalidAttachedTrytes gets returned for invalid attached transaction trytes.
	ErrInvalidAttachedTrytes = errors.New("invalid attached trytes")
	// ErrInvalidTransactionHash gets returned for invalid transaction hash parameters.
	ErrInvalidTransactionHash = errors.New("invalid transaction hash")
	// ErrInvalidTailTransaction gets returned for invalid tail transaction hashes.
	ErrInvalidTailTransaction = errors.New("invalid tail transaction")
	// ErrInvalidTransfer gets returned for invalid transfer parameters.
	ErrInvalidTransfer = errors.New("invalid transfer object")
	// ErrInvalidTrunkTransaction gets returned for invalid trunk transaction parameters.
	ErrInvalidTrunkTransaction = errors.New("invalid trunk transaction")
	// ErrInvalidReferenceHash gets returned for invalid reference hashes.
	ErrInvalidReferenceHash = errors.New("invalid reference hash")
	// ErrInvalidTrytes gets returned for invalid trytes.
	ErrInvalidTrytes = errors.New("invalid trytes")
	// ErrInvalidByte gets returned for an invalid byte for a to trits conversion (5 packed trits in 1 byte).
	ErrInvalidByte = errors.New("invalid byte")
	// ErrInvalidTrit gets returned for invalid trit.
	ErrInvalidTrit = errors.New("invalid trit")
	// ErrInvalidURI gets returned for invalid URIs.
	ErrInvalidURI = errors.New("invalid uri")
	// ErrInvalidASCIIInput gets returned for invalid ASCII input for to trytes conversion.
	ErrInvalidASCIIInput = errors.New("conversion to trytes requires type of input to be encoded in ascii")
	// ErrInvalidOddLength gets returned for odd trytes length for to ASCII conversion.
	ErrInvalidOddLength = errors.New("conversion from trytes requires length of trytes to be even")
	// ErrInvalidTryteEncodedJSON gets returned for invalid tryte encoded JSON messages.
	ErrInvalidTryteEncodedJSON = errors.New("invalid tryte encoded JSON message")
	// ErrSendingBackToInputs gets returned when a transfer sends back to an Input.
	ErrSendingBackToInputs = errors.New("one of the transaction inputs is used as output")
	// ErrNoRemainderSpecified gets returned when no remainder was specified for certain type of operations.
	ErrNoRemainderSpecified = errors.New("remainder address is needed on a transfer with remainder")
)
