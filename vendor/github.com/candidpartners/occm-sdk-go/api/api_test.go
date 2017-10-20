// Package OCCM API implements OnCommand Cloud Manager API.
package api

import (
  "fmt"
  "testing"

  "github.com/candidpartners/occm-sdk-go/util"
  "github.com/candidpartners/occm-sdk-go/api/client"
  "github.com/candidpartners/occm-sdk-go/api/auth"
  "github.com/candidpartners/occm-sdk-go/api/tenant"
  "github.com/candidpartners/occm-sdk-go/api/workenv"
  "github.com/candidpartners/occm-sdk-go/api/workenv/vsa"
  "github.com/candidpartners/occm-sdk-go/api/workenv/awsha"
)

func TestLogin(t *testing.T) {
  context := &client.Context{
    Host: "18.220.2.22",
  }

  authApi, err := auth.New(context)
  if err != nil {
    t.Fatal(err, "Error creating API")
  }

  tenantApi, err := tenant.New(context)
  if err != nil {
    t.Fatal(err, "Error creating Tenant API")
  }

  workenvApi, err := workenv.New(context)
  if err != nil {
    t.Fatal(err, "Error creating Working Environment API")
  }

  vsaWorkenvApi, err := vsa.New(context)
  if err != nil {
    t.Fatal(err, "Error creating VSA Working Environment API")
  }

  awsHaWorkenvApi, err := awsha.New(context)
  if err != nil {
    t.Fatal(err, "Error creating AWS HA Working Environment API")
  }

  fmt.Println("AUTH: ", authApi)
  fmt.Println("TENANT: ", tenantApi)
  fmt.Println("WORKENV: ", workenvApi)
  fmt.Println("WORKENV VSA: ", vsaWorkenvApi)
  fmt.Println("WORKENV AWSHA: ", awsHaWorkenvApi)

  err = authApi.Login("maciej.miechowicz@candidpartners.com", "KZ54TpFHuGeCyjEx")
  if err != nil {
    t.Fatal(err, "Error logging in")
  }

  // tenants, err := tenantApi.GetTenants()
  // if err != nil {
  //   t.Fatal(err, "Error retrieving tenants")
  // }
  // fmt.Println("TENANTS: ", tenants)

  workenvs, err := workenvApi.GetWorkingEnvironments()
  if err != nil {
    t.Fatal(err, "Error retrieving working environments")
  }
  fmt.Println("ENVS: ", workenvs)

  for _, we := range workenvs.VSA {
    fmt.Println("WORK ENV: ", we)

    // aggreagates
    var aggregates []workenv.AggregateResponse
    var err error
    if we.IsHA {
      aggregates, err = awsHaWorkenvApi.GetAggregates(we.PublicId);
    } else {
      aggregates, err = vsaWorkenvApi.GetAggregates(we.PublicId);
    }

    if err != nil {
      t.Fatal(err, "Error retrieving working environment aggregates")
    }
    fmt.Println("AGGREGATES: ", util.ToString(aggregates))

    // // volumes
    // var volumes []workenv.VolumeResponse
    // if we.IsHA {
    //   // volumes, err = awsHaWorkenvApi.GetVolumes(we.PublicId);
    // } else {
    //   volumes, err = vsaWorkenvApi.GetVolumes(we.PublicId);
    // }
    //
    // if err != nil {
    //   t.Fatal(err, "Error retrieving working environment volumes")
    // }
    // fmt.Println("VOLUMES: ", util.ToString(volumes))

    if !we.IsHA {
      // // VSA
      // workenvId := we.PublicId
      // svmName := we.SvmName
      // aggregateName := aggregates[0].Name
      // name := "sample_api_volume"
      // size := &workenv.Capacity{
      //   Size: 5
      //   Unit: "Gigaz"
      // }

      // quoteReq := &vsa.VSAVolumeQuoteRequest{
      //   WorkingEnvironmentId: workenvId,
      //   SvmName: svmName,
      //   Name: "automated_mount",
      //   Size: &workenv.Capacity{
      //     Size: 1,
      //     Unit: "GB",
      //   },
      //   ThinProvisioning: true,
      //   ProviderVolumeType: "gp2",
      //   VerifyNameUniqueness: true,
      //   SnapshotPolicyName: "default",
      //   Compression: true,
      //   Deduplication: true,
      //   ExportPolicyInfo: &workenv.ExportPolicyInfo{
      //     IPs: []string {"10.0.0.0/20"},
      //     PolicyType: "custom",
      //   },
      // }
      // quoteRes, err := vsaWorkenvApi.QuoteVolume(quoteReq)
      // if err != nil {
      //   t.Fatal(err, "Error quoting volume")
      // }
      //
      // fmt.Println("QUOTE RESPONSE: ", util.ToString(quoteRes))
      //
      // req := &vsa.VSAVolumeCreateRequest{
      //   WorkingEnvironmentId: workenvId,
      //   SvmName: svmName,
      //   AggregateName: quoteRes.AggregateName,
      //   Name: "automated_mount",
      //   Size: &workenv.Capacity{
      //     Size: 1,
      //     Unit: "GB",
      //   },
      //   SnapshotPolicyName: "default",
      //   ThinProvisioning: true,
      //   ProviderVolumeType: "gp2",
      //   VerifyNameUniqueness: true,
      //   Compression: true,
      //   Deduplication: true,
      //   ExportPolicyInfo: &workenv.ExportPolicyInfo{
      //     IPs: []string {"10.0.0.0/20"},
      //     PolicyType: "custom",
      //   },
      //   MaxNumOfDisksApprovedToAdd: quoteRes.NumOfDisks,
      // }
      // err = vsaWorkenvApi.CreateVolume(quoteRes.NewAggregate, req);
      // if err != nil {
      //   t.Fatal(err, "Error creating volume")
      // }

      // diskReq := &workenv.VolumeChangeDiskTypeRequest{
      //   AggregateName: aggregateName,
      //   NumOfDisks: 0,
      //   NewAggregate: false,
      //   NewDiskTypeName: "st1",
      // }
      // err = vsaWorkenvApi.ChangeVolumeDiskType(workenvId, svmName, "automated_mount", diskReq);
      // if err != nil {
      //   t.Fatal(err, "Error creating volume")
      // }

      // cloneReq := &workenv.VolumeCloneRequest{
      //   NewVolumeName: "cloned_vol",
      // }
      // err = vsaWorkenvApi.CloneVolume(workenvId, svmName, "automated_mount", cloneReq);
      // if err != nil {
      //   t.Fatal(err, "Error creating volume")
      // }

      // moveReq := &workenv.VolumeMoveRequest{
      //   TargetAggregateName: "someOtherAgg",
      //   NumOfDisksToAdd: 1,
      //   CreateTargetAggregate: true,
      // }
      // err = vsaWorkenvApi.MoveVolume(workenvId, svmName, "automated_mount", moveReq);
      // if err != nil {
      //   t.Fatal(err, "Error creating volume")
      // }

      // err = vsaWorkenvApi.DeleteVolume(workenvId, svmName, "automated_mount");
      // if err != nil {
      //   t.Fatal(err, "Error creating volume")
      // }

    } else {
        // AWS HA
        workenvId := we.PublicId
        // svmName := we.SvmName

        volRes, err := vsaWorkenvApi.GetVolumes(workenvId)
        if err != nil {
          t.Fatal(err, "Error reading volumes")
        }

        fmt.Println("VOLUMES: ", util.ToString(volRes))


        // aggregateName := aggregates[0].Name
        // name := "sample_api_volume"
        // size := &workenv.Capacity{
        //   Size: 5
        //   Unit: "Gigaz"
        // }

        // quoteReq := &vsa.VSAVolumeQuoteRequest{
        //   WorkingEnvironmentId: workenvId,
        //   SvmName: svmName,
        //   Name: "automated_mount",
        //   Size: &workenv.Capacity{
        //     Size: 1,
        //     Unit: "GB",
        //   },
        //   ThinProvisioning: true,
        //   ProviderVolumeType: "gp2",
        //   VerifyNameUniqueness: true,
        //   SnapshotPolicyName: "default",
        //   Compression: true,
        //   Deduplication: true,
        //   ExportPolicyInfo: &workenv.ExportPolicyInfo{
        //     IPs: []string {"10.0.0.0/20"},
        //     PolicyType: "custom",
        //   },
        // }
        // quoteRes, err := awsHaWorkenvApi.QuoteVolume(quoteReq)
        // if err != nil {
        //   t.Fatal(err, "Error quoting volume")
        // }
        // fmt.Println("\n\n\nQUOTE RESPONSE: ", util.ToString(quoteRes))
        //
        // req := &vsa.VSAVolumeCreateRequest{
        //   WorkingEnvironmentId: workenvId,
        //   SvmName: svmName,
        //   AggregateName: quoteRes.AggregateName,
        //   Name: "automated_mount",
        //   Size: &workenv.Capacity{
        //     Size: 1,
        //     Unit: "GB",
        //   },
        //   SnapshotPolicyName: "default",
        //   ThinProvisioning: true,
        //   ProviderVolumeType: "gp2",
        //   VerifyNameUniqueness: true,
        //   Compression: true,
        //   Deduplication: true,
        //   ExportPolicyInfo: &workenv.ExportPolicyInfo{
        //     IPs: []string {"10.0.0.0/20"},
        //     PolicyType: "custom",
        //   },
        //   MaxNumOfDisksApprovedToAdd: quoteRes.NumOfDisks,
        // }
        // err = awsHaWorkenvApi.CreateVolume(quoteRes.NewAggregate, req);
        // if err != nil {
        //   t.Fatal(err, "Error creating volume")
        // }

        // diskReq := &workenv.VolumeChangeDiskTypeRequest{
        //   AggregateName: aggregateName,
        //   NumOfDisks: 0,
        //   NewAggregate: false,
        //   NewDiskTypeName: "st1",
        // }
        // err = awsHaWorkenvApi.ChangeVolumeDiskType(workenvId, svmName, "automated_mount", diskReq);
        // if err != nil {
        //   t.Fatal(err, "Error creating volume")
        // }

        // cloneReq := &workenv.VolumeCloneRequest{
        //   NewVolumeName: "cloned_vol",
        // }
        // err = awsHaWorkenvApi.CloneVolume(workenvId, svmName, "automated_mount", cloneReq);
        // if err != nil {
        //   t.Fatal(err, "Error creating volume")
        // }

        // moveReq := &workenv.VolumeMoveRequest{
        //   TargetAggregateName: "someOtherAgg",
        //   NumOfDisksToAdd: 1,
        //   CreateTargetAggregate: true,
        // }
        // err = awsHaWorkenvApi.MoveVolume(workenvId, svmName, "automated_mount", moveReq);
        // if err != nil {
        //   t.Fatal(err, "Error creating volume")
        // }

        // err = awsHaWorkenvApi.DeleteVolume(workenvId, svmName, "automated_mount");
        // if err != nil {
        //   t.Fatal(err, "Error creating volume")
        // }
    }
  }

}
