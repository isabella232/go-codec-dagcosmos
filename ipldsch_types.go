package dagcosmos

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
)

var _ ipld.Node = nil // suppress errors when this dependency is not referenced
// Type is a struct embeding a NodePrototype/Type for every Node implementation in this package.
// One of its major uses is to start the construction of a value.
// You can use it like this:
//
// 		dagcosmos.Type.YourTypeName.NewBuilder().BeginMap() //...
//
// and:
//
// 		dagcosmos.Type.OtherTypeName.NewBuilder().AssignString("x") // ...
//
var Type typeSlab

type typeSlab struct {
	Address                         _Address__Prototype
	Address__Repr                   _Address__ReprPrototype
	Aunts                           _Aunts__Prototype
	Aunts__Repr                     _Aunts__ReprPrototype
	Block                           _Block__Prototype
	Block__Repr                     _Block__ReprPrototype
	BlockID                         _BlockID__Prototype
	BlockID__Repr                   _BlockID__ReprPrototype
	BlockIDFlag                     _BlockIDFlag__Prototype
	BlockIDFlag__Repr               _BlockIDFlag__ReprPrototype
	Bytes                           _Bytes__Prototype
	Bytes__Repr                     _Bytes__ReprPrototype
	Commit                          _Commit__Prototype
	Commit__Repr                    _Commit__ReprPrototype
	CommitSig                       _CommitSig__Prototype
	CommitSig__Repr                 _CommitSig__ReprPrototype
	Data                            _Data__Prototype
	Data__Repr                      _Data__ReprPrototype
	DuplicateVoteEvidence           _DuplicateVoteEvidence__Prototype
	DuplicateVoteEvidence__Repr     _DuplicateVoteEvidence__ReprPrototype
	Duration                        _Duration__Prototype
	Duration__Repr                  _Duration__ReprPrototype
	Evidence                        _Evidence__Prototype
	Evidence__Repr                  _Evidence__ReprPrototype
	EvidenceData                    _EvidenceData__Prototype
	EvidenceData__Repr              _EvidenceData__ReprPrototype
	EvidenceList                    _EvidenceList__Prototype
	EvidenceList__Repr              _EvidenceList__ReprPrototype
	Hash                            _Hash__Prototype
	Hash__Repr                      _Hash__ReprPrototype
	HashedParams                    _HashedParams__Prototype
	HashedParams__Repr              _HashedParams__ReprPrototype
	Header                          _Header__Prototype
	Header__Repr                    _Header__ReprPrototype
	HexBytes                        _HexBytes__Prototype
	HexBytes__Repr                  _HexBytes__ReprPrototype
	IAVLInnerNode                   _IAVLInnerNode__Prototype
	IAVLInnerNode__Repr             _IAVLInnerNode__ReprPrototype
	IAVLLeafNode                    _IAVLLeafNode__Prototype
	IAVLLeafNode__Repr              _IAVLLeafNode__ReprPrototype
	IAVLNode                        _IAVLNode__Prototype
	IAVLNode__Repr                  _IAVLNode__ReprPrototype
	Int                             _Int__Prototype
	Int__Repr                       _Int__ReprPrototype
	LightBlock                      _LightBlock__Prototype
	LightBlock__Repr                _LightBlock__ReprPrototype
	LightClientAttackEvidence       _LightClientAttackEvidence__Prototype
	LightClientAttackEvidence__Repr _LightClientAttackEvidence__ReprPrototype
	Link                            _Link__Prototype
	Link__Repr                      _Link__ReprPrototype
	MerkleTreeInnerNode             _MerkleTreeInnerNode__Prototype
	MerkleTreeInnerNode__Repr       _MerkleTreeInnerNode__ReprPrototype
	MerkleTreeLeafNode              _MerkleTreeLeafNode__Prototype
	MerkleTreeLeafNode__Repr        _MerkleTreeLeafNode__ReprPrototype
	MerkleTreeNode                  _MerkleTreeNode__Prototype
	MerkleTreeNode__Repr            _MerkleTreeNode__ReprPrototype
	Part                            _Part__Prototype
	Part__Repr                      _Part__ReprPrototype
	PartSet                         _PartSet__Prototype
	PartSet__Repr                   _PartSet__ReprPrototype
	PartSetHeader                   _PartSetHeader__Prototype
	PartSetHeader__Repr             _PartSetHeader__ReprPrototype
	PrivKey                         _PrivKey__Prototype
	PrivKey__Repr                   _PrivKey__ReprPrototype
	Proof                           _Proof__Prototype
	Proof__Repr                     _Proof__ReprPrototype
	Proposal                        _Proposal__Prototype
	Proposal__Repr                  _Proposal__ReprPrototype
	PubKey                          _PubKey__Prototype
	PubKey__Repr                    _PubKey__ReprPrototype
	ResponseDeliverTx               _ResponseDeliverTx__Prototype
	ResponseDeliverTx__Repr         _ResponseDeliverTx__ReprPrototype
	SMTInnerNode                    _SMTInnerNode__Prototype
	SMTInnerNode__Repr              _SMTInnerNode__ReprPrototype
	SMTLeafNode                     _SMTLeafNode__Prototype
	SMTLeafNode__Repr               _SMTLeafNode__ReprPrototype
	SMTNode                         _SMTNode__Prototype
	SMTNode__Repr                   _SMTNode__ReprPrototype
	Signature                       _Signature__Prototype
	Signature__Repr                 _Signature__ReprPrototype
	Signatures                      _Signatures__Prototype
	Signatures__Repr                _Signatures__ReprPrototype
	SignedHeader                    _SignedHeader__Prototype
	SignedHeader__Repr              _SignedHeader__ReprPrototype
	SignedMsgType                   _SignedMsgType__Prototype
	SignedMsgType__Repr             _SignedMsgType__ReprPrototype
	SimpleValidator                 _SimpleValidator__Prototype
	SimpleValidator__Repr           _SimpleValidator__ReprPrototype
	String                          _String__Prototype
	String__Repr                    _String__ReprPrototype
	Time                            _Time__Prototype
	Time__Repr                      _Time__ReprPrototype
	Tx                              _Tx__Prototype
	Tx__Repr                        _Tx__ReprPrototype
	Txs                             _Txs__Prototype
	Txs__Repr                       _Txs__ReprPrototype
	Uint                            _Uint__Prototype
	Uint__Repr                      _Uint__ReprPrototype
	Validator                       _Validator__Prototype
	Validator__Repr                 _Validator__ReprPrototype
	ValidatorSet                    _ValidatorSet__Prototype
	ValidatorSet__Repr              _ValidatorSet__ReprPrototype
	Validators                      _Validators__Prototype
	Validators__Repr                _Validators__ReprPrototype
	Value                           _Value__Prototype
	Value__Repr                     _Value__ReprPrototype
	Version                         _Version__Prototype
	Version__Repr                   _Version__ReprPrototype
	Vote                            _Vote__Prototype
	Vote__Repr                      _Vote__ReprPrototype
}

// --- type definitions follow ---

// Address matches the IPLD Schema type "Address".  It has bytes kind.
type Address = *_Address
type _Address struct{ x []byte }

// Aunts matches the IPLD Schema type "Aunts".  It has list kind.
type Aunts = *_Aunts
type _Aunts struct {
	x []_Hash
}

// Block matches the IPLD Schema type "Block".  It has Struct type-kind, and may be interrogated like map kind.
type Block = *_Block
type _Block struct {
	Header     _Header
	Data       _Data
	Evidence   _EvidenceData
	LastCommit _Commit
}

// BlockID matches the IPLD Schema type "BlockID".  It has Struct type-kind, and may be interrogated like map kind.
type BlockID = *_BlockID
type _BlockID struct {
	Hash          _Link
	PartSetHeader _PartSetHeader
}

// BlockIDFlag matches the IPLD Schema type "BlockIDFlag".  It has int kind.
type BlockIDFlag = *_BlockIDFlag
type _BlockIDFlag struct{ x int64 }

// Bytes matches the IPLD Schema type "Bytes".  It has bytes kind.
type Bytes = *_Bytes
type _Bytes struct{ x []byte }

// Commit matches the IPLD Schema type "Commit".  It has Struct type-kind, and may be interrogated like map kind.
type Commit = *_Commit
type _Commit struct {
	Height     _Int
	Round      _Int
	BlockID    _BlockID
	Signatures _Signatures
}

// CommitSig matches the IPLD Schema type "CommitSig".  It has Struct type-kind, and may be interrogated like map kind.
type CommitSig = *_CommitSig
type _CommitSig struct {
	BlockIDFlag      _BlockIDFlag
	ValidatorAddress _Address
	Timestamp        _Time
	Signature        _Signature
}

// Data matches the IPLD Schema type "Data".  It has Struct type-kind, and may be interrogated like map kind.
type Data = *_Data
type _Data struct {
	Txs _Txs
}

// DuplicateVoteEvidence matches the IPLD Schema type "DuplicateVoteEvidence".  It has Struct type-kind, and may be interrogated like map kind.
type DuplicateVoteEvidence = *_DuplicateVoteEvidence
type _DuplicateVoteEvidence struct {
	VoteA            _Vote
	VoteB            _Vote
	TotalVotingPower _Int
	ValidatorPower   _Int
	Timestamp        _Time
}

// Duration matches the IPLD Schema type "Duration".  It has bytes kind.
type Duration = *_Duration
type _Duration struct{ x []byte }

// Evidence matches the IPLD Schema type "Evidence".
// Evidence has Union typekind, which means its data model behaviors are that of a map kind.
type Evidence = *_Evidence
type _Evidence struct {
	tag uint
	x1  _DuplicateVoteEvidence
	x2  _LightClientAttackEvidence
}
type _Evidence__iface interface {
	_Evidence__member()
}

func (_DuplicateVoteEvidence) _Evidence__member()     {}
func (_LightClientAttackEvidence) _Evidence__member() {}

// EvidenceData matches the IPLD Schema type "EvidenceData".  It has Struct type-kind, and may be interrogated like map kind.
type EvidenceData = *_EvidenceData
type _EvidenceData struct {
	Evidence _EvidenceList
}

// EvidenceList matches the IPLD Schema type "EvidenceList".  It has list kind.
type EvidenceList = *_EvidenceList
type _EvidenceList struct {
	x []_Evidence
}

// Hash matches the IPLD Schema type "Hash".  It has bytes kind.
type Hash = *_Hash
type _Hash struct{ x []byte }

// HashedParams matches the IPLD Schema type "HashedParams".  It has Struct type-kind, and may be interrogated like map kind.
type HashedParams = *_HashedParams
type _HashedParams struct {
	BlockMaxBytes _Int
	BlockMaxGas   _Int
}

// Header matches the IPLD Schema type "Header".  It has Struct type-kind, and may be interrogated like map kind.
type Header = *_Header
type _Header struct {
	Version            _Version
	ChainID            _String
	Height             _Int
	Time               _Time
	LastBlockID        _BlockID
	LastCommitHash     _Link
	DataHash           _Link
	ValidatorsHash     _Link
	NextValidatorsHash _Link
	ConsensusHash      _Link
	AppHash            _Link
	LastResultsHash    _Link
	EvidenceHash       _Link
	ProposerAddress    _Address
}

// HexBytes matches the IPLD Schema type "HexBytes".  It has bytes kind.
type HexBytes = *_HexBytes
type _HexBytes struct{ x []byte }

// IAVLInnerNode matches the IPLD Schema type "IAVLInnerNode".  It has Struct type-kind, and may be interrogated like map kind.
type IAVLInnerNode = *_IAVLInnerNode
type _IAVLInnerNode struct {
	Left    _Link__Maybe
	Right   _Link__Maybe
	Version _Int
	Size    _Int
	Height  _Int
}

// IAVLLeafNode matches the IPLD Schema type "IAVLLeafNode".  It has Struct type-kind, and may be interrogated like map kind.
type IAVLLeafNode = *_IAVLLeafNode
type _IAVLLeafNode struct {
	Key     _Bytes
	Value   _Bytes
	Version _Int
	Size    _Int
	Height  _Int
}

// IAVLNode matches the IPLD Schema type "IAVLNode".
// IAVLNode has Union typekind, which means its data model behaviors are that of a map kind.
type IAVLNode = *_IAVLNode
type _IAVLNode struct {
	tag uint
	x1  _IAVLInnerNode
	x2  _IAVLLeafNode
}
type _IAVLNode__iface interface {
	_IAVLNode__member()
}

func (_IAVLInnerNode) _IAVLNode__member() {}
func (_IAVLLeafNode) _IAVLNode__member()  {}

// Int matches the IPLD Schema type "Int".  It has int kind.
type Int = *_Int
type _Int struct{ x int64 }

// LightBlock matches the IPLD Schema type "LightBlock".  It has Struct type-kind, and may be interrogated like map kind.
type LightBlock = *_LightBlock
type _LightBlock struct {
	SignedHeader _SignedHeader
	ValidatorSet _ValidatorSet
}

// LightClientAttackEvidence matches the IPLD Schema type "LightClientAttackEvidence".  It has Struct type-kind, and may be interrogated like map kind.
type LightClientAttackEvidence = *_LightClientAttackEvidence
type _LightClientAttackEvidence struct {
	ConflictingBlock    _LightBlock
	CommonHeight        _Int
	ByzantineValidators _Validators
	TotalVotingPower    _Int
	Timestamp           _Time
}

// Link matches the IPLD Schema type "Link".  It has link kind.
type Link = *_Link
type _Link struct{ x ipld.Link }

// MerkleTreeInnerNode matches the IPLD Schema type "MerkleTreeInnerNode".  It has Struct type-kind, and may be interrogated like map kind.
type MerkleTreeInnerNode = *_MerkleTreeInnerNode
type _MerkleTreeInnerNode struct {
	Left  _Link__Maybe
	Right _Link__Maybe
}

// MerkleTreeLeafNode matches the IPLD Schema type "MerkleTreeLeafNode".  It has Struct type-kind, and may be interrogated like map kind.
type MerkleTreeLeafNode = *_MerkleTreeLeafNode
type _MerkleTreeLeafNode struct {
	Value _Value
}

// MerkleTreeNode matches the IPLD Schema type "MerkleTreeNode".
// MerkleTreeNode has Union typekind, which means its data model behaviors are that of a map kind.
type MerkleTreeNode = *_MerkleTreeNode
type _MerkleTreeNode struct {
	tag uint
	x1  _MerkleTreeInnerNode
	x2  _MerkleTreeLeafNode
}
type _MerkleTreeNode__iface interface {
	_MerkleTreeNode__member()
}

func (_MerkleTreeInnerNode) _MerkleTreeNode__member() {}
func (_MerkleTreeLeafNode) _MerkleTreeNode__member()  {}

// Part matches the IPLD Schema type "Part".  It has Struct type-kind, and may be interrogated like map kind.
type Part = *_Part
type _Part struct {
	Index _Uint
	Bytes _HexBytes
	Proof _Proof
}

// PartSet matches the IPLD Schema type "PartSet".  It has list kind.
type PartSet = *_PartSet
type _PartSet struct {
	x []_Part
}

// PartSetHeader matches the IPLD Schema type "PartSetHeader".  It has Struct type-kind, and may be interrogated like map kind.
type PartSetHeader = *_PartSetHeader
type _PartSetHeader struct {
	Total _Uint
	Hash  _Link
}

// PrivKey matches the IPLD Schema type "PrivKey".  It has bytes kind.
type PrivKey = *_PrivKey
type _PrivKey struct{ x []byte }

// Proof matches the IPLD Schema type "Proof".  It has Struct type-kind, and may be interrogated like map kind.
type Proof = *_Proof
type _Proof struct {
	Total    _Int
	Index    _Int
	LeafHash _Hash
	Aunts    _Aunts
}

// Proposal matches the IPLD Schema type "Proposal".  It has Struct type-kind, and may be interrogated like map kind.
type Proposal = *_Proposal
type _Proposal struct {
	SMType    _SignedMsgType
	Height    _Int
	Round     _Int
	POLRound  _Int
	BlockID   _BlockID
	Timestamp _Time
	ChainID   _String
}

// PubKey matches the IPLD Schema type "PubKey".  It has bytes kind.
type PubKey = *_PubKey
type _PubKey struct{ x []byte }

// ResponseDeliverTx matches the IPLD Schema type "ResponseDeliverTx".  It has Struct type-kind, and may be interrogated like map kind.
type ResponseDeliverTx = *_ResponseDeliverTx
type _ResponseDeliverTx struct {
	Code      _Uint
	Data      _Bytes
	GasWanted _Int
	GasUsed   _Int
}

// SMTInnerNode matches the IPLD Schema type "SMTInnerNode".  It has Struct type-kind, and may be interrogated like map kind.
type SMTInnerNode = *_SMTInnerNode
type _SMTInnerNode struct {
	Left  _Link__Maybe
	Right _Link__Maybe
}

// SMTLeafNode matches the IPLD Schema type "SMTLeafNode".  It has Struct type-kind, and may be interrogated like map kind.
type SMTLeafNode = *_SMTLeafNode
type _SMTLeafNode struct {
	Path  _Hash
	Value _Hash
}

// SMTNode matches the IPLD Schema type "SMTNode".
// SMTNode has Union typekind, which means its data model behaviors are that of a map kind.
type SMTNode = *_SMTNode
type _SMTNode struct {
	tag uint
	x1  _SMTInnerNode
	x2  _SMTLeafNode
}
type _SMTNode__iface interface {
	_SMTNode__member()
}

func (_SMTInnerNode) _SMTNode__member() {}
func (_SMTLeafNode) _SMTNode__member()  {}

// Signature matches the IPLD Schema type "Signature".  It has bytes kind.
type Signature = *_Signature
type _Signature struct{ x []byte }

// Signatures matches the IPLD Schema type "Signatures".  It has list kind.
type Signatures = *_Signatures
type _Signatures struct {
	x []_CommitSig
}

// SignedHeader matches the IPLD Schema type "SignedHeader".  It has Struct type-kind, and may be interrogated like map kind.
type SignedHeader = *_SignedHeader
type _SignedHeader struct {
	Header _Header
	Commit _Commit
}

// SignedMsgType matches the IPLD Schema type "SignedMsgType".  It has int kind.
type SignedMsgType = *_SignedMsgType
type _SignedMsgType struct{ x int64 }

// SimpleValidator matches the IPLD Schema type "SimpleValidator".  It has Struct type-kind, and may be interrogated like map kind.
type SimpleValidator = *_SimpleValidator
type _SimpleValidator struct {
	PubKey      _PubKey
	VotingPower _Int
}

// String matches the IPLD Schema type "String".  It has string kind.
type String = *_String
type _String struct{ x string }

// Time matches the IPLD Schema type "Time".  It has Struct type-kind, and may be interrogated like map kind.
type Time = *_Time
type _Time struct {
	Seconds     _Int
	Nanoseconds _Int
}

// Tx matches the IPLD Schema type "Tx".  It has list kind.
type Tx = *_Tx
type _Tx struct {
	x []_Bytes
}

// Txs matches the IPLD Schema type "Txs".  It has list kind.
type Txs = *_Txs
type _Txs struct {
	x []_Tx
}

// Uint matches the IPLD Schema type "Uint".  It has bytes kind.
type Uint = *_Uint
type _Uint struct{ x []byte }

// Validator matches the IPLD Schema type "Validator".  It has Struct type-kind, and may be interrogated like map kind.
type Validator = *_Validator
type _Validator struct {
	Address          _Address
	PubKey           _PubKey
	VotingPower      _Int
	ProposerPriority _Int
}

// ValidatorSet matches the IPLD Schema type "ValidatorSet".  It has Struct type-kind, and may be interrogated like map kind.
type ValidatorSet = *_ValidatorSet
type _ValidatorSet struct {
	Validators _Validators
	Proposer   _Validator
}

// Validators matches the IPLD Schema type "Validators".  It has list kind.
type Validators = *_Validators
type _Validators struct {
	x []_Validator
}

// Value matches the IPLD Schema type "Value".
// Value has Union typekind, which means its data model behaviors are that of a map kind.
type Value = *_Value
type _Value struct {
	tag uint
	x1  _SimpleValidator
	x2  _Evidence
	x3  _Link
	x4  _Part
	x5  _ResponseDeliverTx
	x6  _Bytes
	x7  _CommitSig
}
type _Value__iface interface {
	_Value__member()
}

func (_SimpleValidator) _Value__member()   {}
func (_Evidence) _Value__member()          {}
func (_Link) _Value__member()              {}
func (_Part) _Value__member()              {}
func (_ResponseDeliverTx) _Value__member() {}
func (_Bytes) _Value__member()             {}
func (_CommitSig) _Value__member()         {}

// Version matches the IPLD Schema type "Version".  It has Struct type-kind, and may be interrogated like map kind.
type Version = *_Version
type _Version struct {
	Block _Uint
	App   _Uint
}

// Vote matches the IPLD Schema type "Vote".  It has Struct type-kind, and may be interrogated like map kind.
type Vote = *_Vote
type _Vote struct {
	SMType           _SignedMsgType
	Height           _Int
	Round            _Int
	BlockID          _BlockID
	Timestamp        _Time
	ValidatorAddress _Address
	ValidatorIndex   _Int
	Signature        _Signature
}
