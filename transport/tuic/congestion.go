package tuic

import (
	"context"
	"time"

	"github.com/sagernet/quic-go"
	"github.com/sagernet/sing-box/transport/tuic/congestion"
	"github.com/sagernet/sing/common/ntp"
)

func setCongestion(ctx context.Context, connection quic.Connection, congestionName string) {
	timeFunc := ntp.TimeFuncFromContext(ctx)
	if timeFunc == nil {
		timeFunc = time.Now
	}
	switch congestionName {
	case "cubic":
		connection.SetCongestionControl(
			congestion.NewCubicSender(
				congestion.DefaultClock{TimeFunc: timeFunc},
				congestion.GetInitialPacketSize(connection.RemoteAddr()),
				false,
				nil,
			),
		)
	case "new_reno":
		connection.SetCongestionControl(
			congestion.NewCubicSender(
				congestion.DefaultClock{TimeFunc: timeFunc},
				congestion.GetInitialPacketSize(connection.RemoteAddr()),
				true,
				nil,
			),
		)
	case "bbr":
		connection.SetCongestionControl(
			congestion.NewBBRSender(
				congestion.DefaultClock{},
				congestion.GetInitialPacketSize(connection.RemoteAddr()),
<<<<<<< HEAD
				10*congestion.InitialMaxDatagramSize,
=======
				congestion.InitialCongestionWindow*congestion.InitialMaxDatagramSize,
>>>>>>> 0762b71852a168005a3f133e42dc095278fc607a
				congestion.DefaultBBRMaxCongestionWindow*congestion.InitialMaxDatagramSize,
			),
		)
	}
}
