package ipmi

import (
	"fmt"
	"net"
)

func Test() {
	s := NewSimulator(net.UDPAddr{})
	resp := s.reserveRepository(nil)
	reserve, _ := resp.(*ReserveRepositoryResponse)

	err := s.Run()
	client, err := NewClient(s.NewConnection())
	err = client.Open()

	r1 := &SDRFullSensor{}
	r1.Recordid = 5
	r1.Rtype = SDR_RECORD_TYPE_FULL_SENSOR
	r1.SDRVersion = 0x51
	r1.Deviceid = "Fan 5"
	r1.Unit = 0x0
	r1.SensorNumber = 0x04
	r1.SensorType = SDR_SENSOR_TYPECODES_FAN
	r1.BaseUnit = 0x12
	r1.SetMBExp(63, 0, 0, 0)
	r1.ReadingType = SENSOR_READTYPE_THREADHOLD
	data1, _ := r1.MarshalBinary()

	response := &GetSDRCommandResponse{}
	response.CompletionCode = CommandCompleted
	response.NextRecordID = 0xffff

	s.SetHandler(NetworkFunctionStorge, CommandGetSDR, func(m *Message) Response {
		request := &GetSDRCommandRequest{}
		if err := m.Request(request); err != nil {
			return err
		}
		response.ReadData = data1[request.OffsetIntoRecord : request.OffsetIntoRecord+request.ByteToRead]
		return response
	})

	res_senReading := &GetSensorReadingResponse{}
	res_senReading.CompletionCode = CommandCompleted
	res_senReading.SensorReading = 0x2a
	res_senReading.ReadingAvail = 0xc0
	res_senReading.Data1 = 0xc0
	res_senReading.Data2 = 0x00
	s.SetHandler(NetworkFunctionSensorEvent, CommandGetSensorReading, func(m *Message) Response {
		return res_senReading
	})

	sdrSensorInfoList, err := client.GetSensorList(reserve.ReservationId)

	if err == nil {
		if len(sdrSensorInfoList) >= 1 {
			fmt.Println("Fan", sdrSensorInfoList[0].SensorType)
			fmt.Println(float64(2646), sdrSensorInfoList[0].Value)
			fmt.Println("Fan 5", sdrSensorInfoList[0].DeviceId)
			fmt.Println("RPM", sdrSensorInfoList[0].BaseUnit)
			fmt.Println(true, sdrSensorInfoList[0].Avail)
		}
	}
}
