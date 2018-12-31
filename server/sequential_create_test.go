package server_test

import (
	cdsmock "cds.ikigai.net/cabinet.v1/mock"
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"context"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
	"valencex.com/dev/testutil"
)

type testSeqCreate struct{
	uuid string
	seqType string
	seqId uint32
}

func TestServerSequentialCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := cdsmock.NewMockCDSCabinetClient(ctrl)
	tData := []testSeqCreate{
		{uuid: "55DE0AAA-9A6E-4846-88C5-7C2373E5FBE9", seqType: "mock", seqId: 1},
		{uuid: "57A59A9E-5DFB-4926-B0A0-FFE8136E7373", seqType: "mock", seqId: 2},
		{uuid: "F28C14B7-865C-4F58-A9A8-8295E6FDCD60", seqType: "mock", seqId: 3},
		{uuid: "B2D8D095-0DD7-4513-A5B9-C6ECCCECAA0C", seqType: "mock", seqId: 4},
		{uuid: "9A61E569-BCCE-4CDD-ADA5-F33AC432F9E0", seqType: "mock", seqId: 5},
		{uuid: "4F336CC9-F22D-413A-8F89-A2A629A40BB9", seqType: "mock", seqId: 6},
	}

	gCalls := make([]*gomock.Call, len(tData))

	for i := range tData{
		gCalls[i] = m.EXPECT().SequentialCreate(
			gomock.Any(), &rpcMsg{msg: &pb.Sequential{Uuid: tData[i].uuid, Type: tData[i].seqType}},
		).Return(&pb.Sequential{Seqid: tData[i].seqId, Uuid: tData[i].uuid}, nil)

		if i > 0{
			gCalls[i].After(gCalls[i-1])
		}
	}

	testSequentialCreate(t, m, tData)
}

func testSequentialCreate(t *testing.T, client *cdsmock.MockCDSCabinetClient, data []testSeqCreate) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tx := testutil.NewTestRunner(t)

	for i := range data {
		r, err := client.SequentialCreate(ctx, &pb.Sequential{Uuid: data[i].uuid, Type: data[i].seqType})

		tx.AssertNilError(err, "SequentialCreate")
		tx.AssertEqualString(r.Uuid, data[i].uuid, "Sequential.uuid")
		tx.AssertEqualUInt32(r.Seqid, data[i].seqId, "Sequential.seqId")

		t.Log("SequentialCreate: ", r)
	}
}
