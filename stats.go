package segment

type CollectionStats interface {

	// TotalDocumentCount returns the number of documents, regardless of whether or not
	// they have any terms for this field
	TotalDocumentCount() int

	// DocumentCount returns the number of documents with at least one term for this field
	DocumentCount() int

	// SumTotalTermFrequency returns to total number of tokens across all documents
	SumTotalTermFrequency() int

	// SumDocumentFrequency returns the sum of all posting list entries for this field
	// SumDocumentFrequency() int

	Merge(CollectionStats)
}

type TermStats interface {

	// DocumentFrequency returns the number of documents using this term
	DocumentFrequency() int

	// TotalTermFrequency returns the total number of occurrences of this term
	// TotalTermFrequency() int
}
