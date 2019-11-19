USE tipster;

CREATE TABLE Tip (
    TipId int NOT NULL AUTO_INCREMENT,
    Bookie varchar(255),
    Analysis varchar(255),
    Rate float,
	Stake int,
	Pick varchar(255),
	PRIMARY Key(TipId)
);


CREATE TABLE Rivals (
    RivalId int NOT NULL AUTO_INCREMENT,
    RivalName varchar(255),
	TipId int,
	PRIMARY KEY(RivalId),
    FOREIGN KEY (TipId) REFERENCES Tip(TipId)
);