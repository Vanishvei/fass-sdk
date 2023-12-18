# Taocloud block storage api sdk (fass-sdk)

### 操作对象

| 操作               | 请求对象                     | 参数对象                     | 返回对象                    |
|------------------|--------------------------|--------------------------|-------------------------|
| 获取存储池信息          | RetrievePool             | RetrievePool             | RetrievePool            |
| 创建账户             | CreateAccount            | CreateAccount            | CreateAccount           |
| 获取帐户列表           | ListAccount              | ListAccount              | ListAccount             |
| 获取帐户详情           | RetrieveAccount          | RetrieveAccount          | RetrieveAccount         |
| 删除账户             | DeleteAccount            | DeleteAccount            | error                   |
| 添加主机到主机组         | AddHostToHostGroup       | AddHostToHostGroup       | AddHostToHostGroup      |
| 获取主机组关联的subsys列表 | ListSubsysOfHostGroup    | ListSubsysOfHostGroup    | ListSubsysOfHostGroup   |
| 获取主机组列表          | ListHostGroup            | ListHostGroup            | ListHostGroup           |
| 从主机组移除主机         | RemoveHostFromHostGroup  | RemoveHostFromHostGroup  | RemoveHostFromHostGroup |
| 删除主机组            | DeleteHostGroup          | DeleteHostGroup          | error                   |
| 创建subsys         | CreateSubsys             | CreateSubsys             | CreateSubsys            |
| 获取subsys详情       | RetrieveSubsys           | RetrieveSubsys           | RetrieveSubsys          |
| 从卷创建subsys       | CreateSubsysFromVolume   | CreateSubsysFromVolume   | CreateSubsys            |
| 从快照创建subsys      | CreateSubsysFromSnapshot | CreateSubsysFromSnapshot | CreateSubsys            |
| 获取subsys列表       | ListSubsys               | ListSubsys               | ListSubsys              |
| subsys导出协议       | ExportSubsys             | ExportSubsys             | error                   |
| subsys取消协议导出     | UnexportSubsys           | UnexportSubsys           | error                   |
| subsys绑定主机组      | SubsysBindHostGroup      | SubsysBindHostGroup      | error                   |
| subsys取消绑定的主机组   | SubsysUnbindHostGroup    | SubsysUnbindHostGroup    | error                   |
| 获取subsys绑定的主机组   | RetrieveSubsysHostGroup  | RetrieveSubsysHostGroup  | RetrieveSubsysHostGroup |
| 设置Chap           | SetSubsysChap            | SetSubsysChap            | error                   |
| 移除Chap           | RemoveSubsysChap         | RemoveSubsysChap         | error                   |
| 获取subsys的chap详情  | RetrieveSubsysChap       | RetrieveSubsysChap       | RetrieveSubsysChap      |
| subsys添加VLAN项    | SubsysAddVLAN            | SubsysAddVLAN            | error                   |
| 获取subsys的VLAN详情  | RetrieveSubsysVLAN       | RetrieveSubsysVLAN       | RetrieveSubsysVLAN      |
| subsys移除VLAN项    | SubsysRemoveVLAN         | SubsysRemoveVLAN         | error                   |
| subsys删除VLAN     | DeleteSubsysVLAN         | DeleteSubsysVLAN         | error                   |
| 删除subsys         | DeleteSubsys             | DeleteSubsys             | error                   |
| 创建快照             | CreateSnapshot           | CreateSnapshot           | CreateSnapshot          |
| 获取快照详情           | RetrieveSnapshot         | RetrieveSnapshot         | RetrieveSnapshot        |
| 获取卷的快照列表         | ListSnapshot             | ListSnapshot             | ListSnapshot            |
| 回滚快照             | RevertSnapshot           | RevertSnapshot           | error                   |
| 删除快照             | DeleteSnapshot           | DeleteSnapshot           | error                   |
| 获取卷列表            | ListVolume               | ListVolume               | ListVolume              |
| 获取卷详情            | RetrieveVolume           | RetrieveVolume           | RetrieveVolume          |
| 卷扩容              | ExpandVolume             | ExpandVolume             | error                   |
| 卷设置流控            | SetQosOfVolume           | SetQosOfVolume           | error                   |
| 分离卷              | FlattenVolume            | FlattenVolume            | FlattenVolume           |
| 获取卷分离进度          | GetFlattenVolumeProgress | GetFlattenVolumeProgress | FlattenVolumeProgress   |
| 停止卷分离            | StopFlattenVolume        | StopFlattenVolume        | error                   |

-------

### 示例

**初始化配置**
```
var (
    port           = 8000
    backoff        = 0
    retryCount     = 3
    readTimeout    = 30
    connectTimeout = 1
    endpointList   = []string{"172.18.165.200", "172.18.165.201", "172.18.165.202"}
)

fassSDK.InitConfig(
    &endpointList,
    &port,
    &readTimeout,
    &connectTimeout,
    &backoff,
    &retryCount,
)
```

**获取存储池信息**
```
package pool

import (    
    "fmt"
    
    "github.com/google/uuid"
    
    fassSDK "github.com/Vanishvei/fass-sdk"
    parameters "github.com/Vanishvei/fass-sdk-parameters"
    _ "github.com/Vanishvei/fass-sdk-responses"
    _ "github.com/Vanishvei/fass-sdk/test"
)

func RetrievePool() {
    parameter := parameters.RetrievePool{}
    parameter.SetPoolName(poolName)
    
    resp, err := fassSDK.RetrievePool(&parameter, uuid.New().String())
    if err != nil {
        fmt.Printf("%s", err.Error())
        return
    }        
	
    fmt.Printf("%s\n", resp.Name)
    fmt.Printf("%v", resp)
}

```

**创建subsys**
```
func CreateSubsys() {
    parameter := parameters.CreateSubsys{}
    parameter.SetName("subsys001")
    parameter.SetPoolName("fast_pool")
    parameter.SetVolumeName("volume001")
    parameter.SetIops(2000)
    parameter.SetBpsMB(1000)
    parameter.SetCapacityGB(10)
    parameter.SetStripeWidth4Shift128k()
    parameter.SetSectorSize4096()
    parameter.SetFormatROW()
    parameter.EnableISCSI()

    resp, err := fassSDK.CreateSubsys(&parameter, uuid.New().String())
    if err != nil {
        fmt.Printf("%s", err.Error())
        return
    }

    fmt.Printf("%s\n", resp.SubsysName)
    fmt.Printf("%v\n", resp)
}
```

**获取subsys详细信息**
```
func RetrieveSubsys() {
    parameter := parameters.RetrieveSubsys{}
    parameter.SetSubsysName(subsysName)

    resp, err := fassSDK.RetrieveSubsys(&parameter, uuid.New().String())
    if err != nil {
        fmt.Printf("%s", err.Error())
        return 
    }
	
    fmt.Printf("%v\n", resp)
}
```

**从快照创建subsys**
```
func CreateSubsysFromSnapshot() {
    parameter := parameters.CreateSubsysFromSnapshot{}
    parameter.SetName("subsys002")
    parameter.SetPoolName("fast_pool")
    parameter.SetSrcVolumeName("volume001")
    parameter.SetSrcSnapshotName("snapshot001")
    parameter.InheritQos()
    parameter.EnableISCSI()

    resp, err := fassSDK.CreateSubsysFromSnapshot(&parameter, requestId)
    if err != nil {
        fmt.Printf("%s\n", err.Error())
        return
    }
    
    fmt.Printf("%v\n", resp)
}
```

**subsys导出协议**
```
func ExportSubsys() {
    parameter := parameters.ExportSubsys{}
    parameter.SetSubsysName("subsys001")
    parameter.ExportISCSI()

    err := fassSDK.ExportSubsys(&parameter, uuid.New().String())
    if err != nil {
    	fmt.Printf("%s", err.Error())
    	return
    }
}
```

**subsys取消导出协议**
```
func UnexportSubsys() {
	parameter := parameters.UnexportSubsys{}
	parameter.SetSubsysName("subsys001")
	parameter.UnexportISCSI()

	err := fassSDK.UnexportSubsys(&parameter, uuid.New().String())
    if err != nil {
    	fmt.Printf("%s", err.Error())
    	return
    }
}
```


**删除subsys**
```
func DeleteSubsys() error {
    parameter := parameters.DeleteSubsys{}
    parameter.SetSubsysName("subsys001")
    parameter.DeleteVolume()
    parameter.ForceDelete()

    err := fassSDK.DeleteSubsys(&parameter, uuid.New().String())
    if err != nil {
        fmt.Printf("%s\n", err.Error())
    }
}

```

**创建快照**
```
func CreateSnapshot() {
    parameter := parameters.CreateSnapshot{}
    parameter.SetVolumeName("volume001)
    parameter.SetSnapshotName("snapshot001")
	
    resp, err := fassSDK.CreateSnapshot(&parameter, uuid.New().String())
    if err != nil {
        fmt.Printf("%s", err.Error())
        return
    }
	
    fmt.Printf("%v\n", resp)
}
```

**删除快照**
```
func DeleteSnapshot() {
    parameter := parameters.DeleteSnapshot{}
    parameter.SetVolumeName("volume001")
    parameter.SetSnapshotName("snapshot001")
	
    err := fassSDK.DeleteSnapshot(&parameter, uuid.New().String())
    if err != nil {
        fmt.Printf("%s\n", err.Error())
    }
}
```

**设置卷的流控**
```
func SetQosOfVolume() {
    parameter := parameters.SetQosOfVolume{}
    parameter.SetVolumeName("volume001")
    parameter.SetBpsMB(100)
    parameter.SetIops(1000)
    
    err := fassSDK.SetQosOfVolume(&parameter, uuid.New().String())
    if err != nil) {
    	fmt.Printf("%s\n", err.Error())
    }
}
```

**卷扩容**
```
func ExpandVolume() {
    parameter := parameters.ExpandVolume{}
    parameter.SetVolumeName("volume001")
    parameter.SetNewCapacityGB(20)

    err := fassSDK.ExpandVolume(&parameter, uuid.New().String())
    if err != nil {
        fmt.Printf("%s", err.Error())
        return
    }
}
```

**分离卷(仅对使用快照创建出来的卷和使用卷创建出来的卷可用)**
```
func flattenVolume(volumeName string) (*responses.FlattenVolume, error) {
    parameter := parameters.FlattenVolume{}
    parameter.SetVolumeName("volume002")

    data, err := fassSDK.FlattenVolume(&parameter, uuid.New().String())
    if err != nil {
        fmt.Printf("%s\n", err.Error())
    }
    
    fmt.Printf("task_id: %s\n", data.TaskId)    
```

**查询卷分离进度**
```
func FlattenVolumeProgress() {
    parameter := parameters.GetFlattenVolumeProgress{}
    parameter.SetTaskId(taskId)

    data, err := fassSDK.FlattenVolumeProgress(&parameter, uuid.New().String())
    if err != nil {
    	fmt.Printf("%s", err.Error())
    	return
    }
    
    if data.IsDone() {
    	fmt.Printf("volume flatten complete\n")
    	return
    }
    
    fmt.Printf("%s\n". data.Status)
}
```

**停止卷分离**
```
func StopFlattenVolume() {
    data, err := flattenVolume("volume002")
    if !reflect.DeepEqual(err, nil) {
    	fmt.Printf("%s", err.Error())
    	return 
    }

    parameter := parameters.StopFlattenVolume{}
    parameter.SetTaskId(data.TaskId)

    err := fassSDK.StopFlattenVolume(&parameter, uuid.New().String())
    if err != nil {
        fmt.Printf("%s", err.Error())
    	return 
    }
}
```
----

### 错误码

| 错误码    | 错误类型                      | 
|--------|---------------------------|
| 100001 | 	OperationNotSupport      |
| 100002 | 	OperationNotPermit       |
| 100003 | 	OperationFailed          |
| 100004 | 	OperationAborted         |
| 100005 | 	OperationCanceled        |
| 100006 | 	OperationTimeout         |
| 100007 | 	ServiceUnavailable       |
| 100008 | 	SuzakuResultParseFailed  |
| 100009 | 	SlowDown                 |
| 100010 | 	RpcDiscoveryFailed       |
| 100011 | 	RpcRequestFailed         |
| 130000 | 	InvalidRecoveryQOS       |
| 130001 | 	InvalidBalanceQOS        |
| 200001 | 	TaskNotExist             |
| 200002 | 	TaskCannotStop           |
| 300001 | 	InvalidRequest           |
| 300002 | 	InvalidRequestBody       |
| 300003 | 	MissingContentLength     |
| 300004 | 	MissingRequestBody       |
| 300005 | 	MissingRequestHeader     |
| 300006 | 	InvalidArgument          |
| 300007 | 	MethodNotAllowed         |
| 400001 | 	StorageSpaceInsufficient |
| 400002 | 	PoolAlreadyExists        |
| 400003 | 	PoolNotExist             |
| 400004 | 	PoolTooMany              |
| 400005 | 	PoolIsNotEmpty           |
| 400006 | 	PoolStatusConfirm        |
| 500001 | 	SubsysAlreadyExists      |
| 500002 | 	SubsysNotExist           |
| 500003 | 	SubsysInUse              |
| 500004 | 	SubsysTooMany            |
| 500005 | 	SubsysStatusConfirm      |
| 500006 | 	SubsysNotBindHostGroup   |
| 500007 | 	SubsysAlreadyHostGroup   |
| 500008 | 	SubsysNotSetChap         |
| 500009 | 	SubsysBusy               |
| 510001 | 	VLANDuplicate            |
| 510002 | 	VLANNotExists            |
| 600001 | 	VolumeAlreadyExists      |
| 600002 | 	VolumeNotExist           |
| 600003 | 	VolumeInUse              |
| 600004 | 	VolumeTooMany            |
| 600005 | 	VolumeCannotShrink       |
| 600006 | 	VolumeOnFlatten          |
| 600007 | 	VolumeCapacityToLarge    |
| 600008 | 	VolumeQosTooSmall        |
| 600009 | 	VolumeStatusConfirm      |
| 700001 | 	SnapshotAlreadyExists    |
| 700002 | 	SnapshotNotExist         |
| 700003 | 	SnapshotInUse            |
| 700004 | 	SnapshotTooMany          |
| 700005 | 	SnapshotStatusConfirm    |
| 800001 | 	AccountAlreadyExists     |
| 800002 | 	AccountNotExists         |
| 801001 | 	ChapIsApply              |
| 810001 | 	HostGroupNotExists       |
| 810002 | 	HostAlreadyExists        |
| 810003 | 	HostNotExists            |
| 810004 | 	HostGroupNotEmpty        |
| 810005 | 	HostBusy                 |
| 810006 | 	HostGroupMapping         |
