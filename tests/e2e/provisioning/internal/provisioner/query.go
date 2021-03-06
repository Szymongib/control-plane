package provisioner

import "fmt"

type queryProvider struct{}

func (qp queryProvider) upgradeRuntime(runtimeID string, config string) string {
	return fmt.Sprintf(`mutation {
	result: upgradeRuntime(id: "%s", config: %s) {
    	%s
  	}
}`, runtimeID, config, operationStatusData())
}

func (qp queryProvider) runtimeStatus(operationID string) string {
	return fmt.Sprintf(`query {
	result: runtimeStatus(id: "%s") {
		%s
	}
}`, operationID, runtimeStatusData())
}

func (qp queryProvider) runtimeOperationStatus(operationID string) string {
	return fmt.Sprintf(`query {
	result: runtimeOperationStatus(id: "%s") {
		%s
	}
}`, operationID, operationStatusData())
}

func runtimeStatusData() string {
	return fmt.Sprintf(`lastOperationStatus { operation state message }
			runtimeConnectionStatus { status }
			runtimeConfiguration { 
				kubeconfig
				clusterConfig { 
					%s
				} 
				kymaConfig { 
					version
				} 
			}`, clusterConfig())
}

func clusterConfig() string {
	return fmt.Sprintf(`
		name
		kubernetesVersion
		volumeSizeGB
		diskType
		machineType
		region
		provider
		seed
		targetSecret
		diskType
		workerCidr
		autoScalerMin
		autoScalerMax
		maxSurge
		maxUnavailable
		providerSpecificConfig {
			%s
		}
`, providerSpecificConfig())
}

func providerSpecificConfig() string {
	return fmt.Sprint(`
		... on GCPProviderConfig { 
			zones 
		} 
		... on AzureProviderConfig {
			vnetCidr
		}
		... on AWSProviderConfig {
			zones
			internalCidr 
			vpcCidr 
			publicCidr
		}  
	`)
}

func operationStatusData() string {
	return `id
			operation 
			state
			message
			runtimeID`
}
