package codec

import "github.com/dfuse-io/bstream"

func init() {
	bstream.GetBlockWriterFactory = bstream.BlockWriterFactoryFunc(BlockWriterFactory)
	bstream.GetBlockReaderFactory = bstream.BlockReaderFactoryFunc(BlockReaderFactory)
	bstream.GetBlockDecoder = bstream.BlockDecoderFunc(BlockDecoder)
	bstream.GetProtocolFirstBlock = 2
}
