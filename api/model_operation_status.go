/*
 * Rosetta
 *
 * <h2>Backstory</h2> Writing reliable blockchain integrations is complicated and time-consuming. The process requires careful analysis of the unique aspects of each blockchain and extensive communication with its developers to understand the best strategies to deploy nodes, recognize deposits, broadcast transactions, etc. Even a minor misunderstanding can lead to downtime, or even worse, incorrect fund attribution. Not to mention, this integration must be continuously modified and tested each time a blockchain team releases new software.  Instead of spending time working on their blockchain, project developers spend countless hours answering similar support questions for each team integrating their blockchain. With their questions answered, each integrating team then writes similar code to interface with the blockchain instead of spending their engineering resources adding support for more blockchain projects or working on unique products and applications.  <h2>Standard for Blockchain Interaction</h2> Rosetta is a new project from Coinbase to standardize the process of deploying and interacting with blockchains. With an explicit specification to adhere to, all parties involved in blockchain development can spend less time figuring out how to integrate with each other and more time working on the novel advances that will push the blockchain ecosystem forward. In practice, this means that any blockchain project that implements the requirements outlined in this specification will enable exchanges, block explorers, and wallets to integrate with much less communication overhead and network-specific work.  <h5>© 2020 Coinbase</h5>
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type OperationStatus struct {

	// The `status` is the network-specific status of the operation.
	Status string `json:"status"`

	// An `Operation` is considered `successful` if the `Operation.Amount` should affect the `Operation.Account`. Some blockchains (like Bitcoin) only include `successful` operations in blocks but other blockchains (like Ethereum) include unsuccessful operations that incur a fee.  To reconcile the computed balance from the stream of `Operations`, it is critical to understand which `Operation.Status` indicate an `Operation` is `successful` and should affect an `Account`.
	Successful bool `json:"successful"`
}
