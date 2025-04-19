/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

SET SESSION innodb_strict_mode = OFF;

DROP DATABASE IF EXISTS `bilgaden_new_new`;
CREATE DATABASE IF NOT EXISTS `bilgaden_new_new` /*!40100 DEFAULT CHARACTER SET latin1 COLLATE latin1_danish_ci */;
USE `bilgaden_new_new`;

DROP TABLE IF EXISTS `koeretoej`;
CREATE TABLE IF NOT EXISTS `koeretoej` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  PRIMARY KEY (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejanvendelse`;
CREATE TABLE IF NOT EXISTS `koeretoejanvendelse` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejAnvendelseNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejAnvendelseNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejAnvendelseBeskrivelse` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `KoeretoejAnvendelseGyldigFra` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejAnvendelseGyldigTil` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejAnvendelseStatus` varchar(40) DEFAULT NULL COMMENT 'ForretningParameterVaerdiStatusType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejAnvendelseNummer`),
  CONSTRAINT `koeretoejanvendelse_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejanvendelsesupplerende`;
CREATE TABLE IF NOT EXISTS `koeretoejanvendelsesupplerende` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejAnvendelseSupplerendeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejAnvendelseSupplerendeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejAnvendelseSupplerendeBeskrivelse` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `KoeretoejAnvendelseSupplerendeGyldigFra` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejAnvendelseSupplerendeGyldigTil` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejAnvendelseSupplerendeStatus` varchar(40) DEFAULT NULL COMMENT 'ForretningParameterVaerdiStatusType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejAnvendelseSupplerendeNummer`),
  CONSTRAINT `koeretoejanvendelsesupplerende_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejart`;
CREATE TABLE IF NOT EXISTS `koeretoejart` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejArtNummer` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejArtNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejArtKraeverForsikring` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejArtBeskrivelse` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `KoeretoejArtGyldigFra` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejArtGyldigTil` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejArtStatus` varchar(40) DEFAULT NULL COMMENT 'ForretningParameterVaerdiStatusType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejArtNummer`),
  CONSTRAINT `koeretoejart_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejblokeringaarsag`;
CREATE TABLE IF NOT EXISTS `koeretoejblokeringaarsag` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejBlokeringAarsagTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejBlokeringAarsagTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejBlokeringAarsagTypeNummer`),
  CONSTRAINT `koeretoejblokeringaarsag_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejbraendstof`;
CREATE TABLE IF NOT EXISTS `koeretoejbraendstof` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejMotorKmPerLiter` decimal(6,1) DEFAULT NULL COMMENT 'TalDecimalType.xsd',
  `KoeretoejMotorKMPerLiterPreCalc` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorBraendstofforbrugMaalt` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorGasforbrug` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorCO2UdslipBeregnet` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMiljoeOplysningCO2Udslip` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejbraendstof_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejco2emissionklasse`;
CREATE TABLE IF NOT EXISTS `koeretoejco2emissionklasse` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `CO2EmissionKlasseNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `CO2EmissionKlasseNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`CO2EmissionKlasseNummer`),
  CONSTRAINT `koeretoejco2emissionklasse_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejdispensationtype`;
CREATE TABLE IF NOT EXISTS `koeretoejdispensationtype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `DispensationTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `DispensationTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejDispensationTypeKommentar` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`DispensationTypeNummer`),
  CONSTRAINT `koeretoejdispensationtype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejdrivkrafttype`;
CREATE TABLE IF NOT EXISTS `koeretoejdrivkrafttype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `DrivkraftTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `DrivkraftTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`DrivkraftTypeNummer`),
  CONSTRAINT `koeretoejdrivkrafttype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejdrivmiddel`;
CREATE TABLE IF NOT EXISTS `koeretoejdrivmiddel` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejMotorBraendselscelle` tinyint(1) DEFAULT NULL COMMENT 'MarkeringType.xsd',
  `KoeretoejMotorPlugInHybrid` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejMotorDrivmiddelPrimaer` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejdrivmiddel_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejelforbrug`;
CREATE TABLE IF NOT EXISTS `koeretoejelforbrug` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejMotorElektriskForbrug` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorElektriskForbrugMaalt` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorElektriskRaekkevidde` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorBatterikapacitet` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejelforbrug_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejfarvetype`;
CREATE TABLE IF NOT EXISTS `koeretoejfarvetype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `FarveTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `FarveTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`FarveTypeNummer`),
  CONSTRAINT `koeretoejfarvetype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejfastkombination`;
CREATE TABLE IF NOT EXISTS `koeretoejfastkombination` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejFastKombinationIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `RegistreringNummerNummer` varchar(10) DEFAULT NULL COMMENT 'RegistreringNummerType.xsd',
  `RegistreringNummerIdent` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejFastKombinationIdent`),
  CONSTRAINT `koeretoejfastkombination_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejgruppe`;
CREATE TABLE IF NOT EXISTS `koeretoejgruppe` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejGruppeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejGruppeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejGruppeNummer`),
  CONSTRAINT `koeretoejgruppe_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejkarrosseritype`;
CREATE TABLE IF NOT EXISTS `koeretoejkarrosseritype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KarrosseriTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KarrosseriTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KarrosseriTypeNummer`),
  CONSTRAINT `koeretoejkarrosseritype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejleasing`;
CREATE TABLE IF NOT EXISTS `koeretoejleasing` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `LeasingMaaneder` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `LeasingNummer` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `LeasingGyldigFra` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `LeasingGyldigTil` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `LeasingReelOphoerDato` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `LeasingKode` varchar(10) DEFAULT NULL COMMENT 'KodeType.xsd',
  `LeasingStatus` varchar(25) DEFAULT NULL COMMENT 'LeasingStatusType.xsd',
  `LeasingBemaerkning` varchar(2000) DEFAULT NULL COMMENT 'Tekst2000Type.xsd',
  `LeasingAendringType` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `LeasingSidstAendret` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejleasing_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejmaalenorm`;
CREATE TABLE IF NOT EXISTS `koeretoejmaalenorm` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejMotorMaaleNormTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejMotorMaaleNormTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejMotorMaaleNormTypeNummer`),
  CONSTRAINT `koeretoejmaalenorm_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejmaerketype`;
CREATE TABLE IF NOT EXISTS `koeretoejmaerketype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejMaerkeTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejMaerkeTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejMaerkeTypeNummer`),
  CONSTRAINT `koeretoejmaerketype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejmiljoeoplysning`;
CREATE TABLE IF NOT EXISTS `koeretoejmiljoeoplysning` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejMiljoeOplysningCO2EmissionKlasseOmklassificeringDato` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejMiljoeOplysningCO2EmissionKlasseOprindelig` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejMiljoeOplysningCO2Udslip` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMiljoeOplysningEftermonteretPartikelfilter` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejMiljoeOplysningEmissionCO` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMiljoeOplysningEmissionHCPlusNOX` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMiljoeOplysningEmissionNOX` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMiljoeOplysningNyttelastvaerdi` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMiljoeOplysningPartikelFilter` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejMiljoeOplysningPartikler` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMiljoeOplysningRoegtaethed` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMiljoeOplysningRoegtaethedOmdrejningstal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejMiljoeOplysningSpecifikCO2Emission` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMiljoeOplysningTungtNulEmissionKoeretoej` tinyint(1) DEFAULT NULL COMMENT 'JaNejType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejmiljoeoplysning_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejmodeltype`;
CREATE TABLE IF NOT EXISTS `koeretoejmodeltype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejModelTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejModelTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejModelTypeNummer`),
  CONSTRAINT `koeretoejmodeltype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejmotor`;
CREATE TABLE IF NOT EXISTS `koeretoejmotor` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejMotorCylinderAntal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejMotorSlagVolumen` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorSlagVolumenIkkeTilgaengelig` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejMotorStoersteEffekt` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorStoersteEffektIkkeTilgaengelig` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejMotorKilometerstand` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejMotorKilometerstandDokumentation` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejMotorKilometerstandIkkeTilgaengelig` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejMotorKmPerLiter` decimal(6,1) DEFAULT NULL COMMENT 'TalDecimalType.xsd',
  `KoeretoejMotorKMPerLiterPreCalc` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorPlugInHybrid` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejMotorMaerkning` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejMotorStandStoej` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorKoerselStoej` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorStandStoejOmdrejningstal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejMotorInnovativTeknik` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejMotorInnovativTeknikAntal` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorElektriskForbrug` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorFuelmode` varchar(10) DEFAULT NULL COMMENT 'FuelmodeTypeType.xsd',
  `KoeretoejMotorGasforbrug` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorElektriskRaekkevidde` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorBatterikapacitet` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorBraendstofforbrugMaalt` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorElektriskForbrugMaalt` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorMaaleNormTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejMotorMaaleNormTypeNummer` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejMotorCO2UdslipBeregnet` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejMotorBraendselscelle` tinyint(1) DEFAULT NULL COMMENT 'MarkeringType.xsd',
  `KoeretoejMotorDrivmiddelPrimaer` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejmotor_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejnormtype`;
CREATE TABLE IF NOT EXISTS `koeretoejnormtype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `NormTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `NormTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`NormTypeNummer`),
  CONSTRAINT `koeretoejnormtype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejoplysning`;
CREATE TABLE IF NOT EXISTS `koeretoejoplysning` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejOplysningOprettetUdFra` varchar(100) DEFAULT NULL COMMENT 'KoeretoejOplysningOprettetUdFraType.xsd',
  `KoeretoejOplysningStatus` varchar(100) DEFAULT NULL COMMENT 'KoeretoejOplysningStatusType.xsd',
  `KoeretoejOplysningStatusDato` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `KoeretoejOplysningFoersteRegistreringDato` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejOplysningStelNummer` varchar(20) DEFAULT NULL COMMENT 'StelNummerType.xsd',
  `KoeretoejOplysningStelNummerAnbringelse` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `KoeretoejOplysningModelAar` int(4) DEFAULT NULL COMMENT 'AArstalType.xsd',
  `KoeretoejOplysningTotalVaegt` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningEgenVaegt` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningKoereklarVaegtMinimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningKoereklarVaegtMaksimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningTekniskTotalVaegt` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningVogntogVaegt` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningAkselAntal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningStoersteAkselTryk` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningSkatteAkselAntal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningSkatteAkselTryk` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningPassagerAntal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningSiddepladserMinimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningSiddepladserMaksimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningStaapladserMinimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningStaapladserMaksimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningTilkoblingMulighed` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejOplysningTilkoblingsvaegtUdenBremser` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningTilkoblingsvaegtMedBremser` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningPaahaengVognTotalVaegt` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningSkammelBelastning` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningSaettevognTilladtAkselTryk` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningMaksimumHastighed` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningFaelgDaek` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `KoeretoejOplysningTilkobletSidevognStelnr` varchar(20) DEFAULT NULL COMMENT 'StelNummerType.xsd',
  `KoeretoejOplysningNCAPTest` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejOplysningVVaerdiLuft` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejOplysningVVaerdiMekanisk` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejOplysningOevrigtUdstyr` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `KoeretoejOplysningKoeretoejstand` varchar(20) DEFAULT NULL COMMENT 'KoeretoejstandType.xsd',
  `KoeretoejOplysning30PctVarevogn` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejOplysningBlokvognAkselType` varchar(40) DEFAULT NULL COMMENT 'AkselTypeType.xsd',
  `KoeretoejOplysningBlokvognHovedboltTryk` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningBlokvognSkammelTryk` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningBlokvognSamletAkselTryk` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningBlokvognMaxVogntog` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningBlokvognBreddeFra` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejOplysningBlokvognKoblingshoejdeFra` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejOplysningBlokvognKoblingslaengdeFra` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejOplysningBlokvognSammenkoblingType` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejOplysningBlokvognTilladeligHastighed` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningBlokvognBreddeTil` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejOplysningBlokvognKoblingshoejdeTil` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejOplysningBlokvognKoblingslaengdeTil` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejOplysningTraekkendeAksler` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejOplysningEgnetTilTaxi` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejOplysningAkselAfstand` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningSporviddenForrest` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningSporviddenBagest` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningTypeAnmeldelseNummer` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejOplysningTypeGodkendelseNummer` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejOplysningEUVariant` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejOplysningEUVersion` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejOplysningKommentar` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `KoeretoejOplysningTypegodkendtKategori` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejOplysningAntalGear` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningAntalDoere` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejOplysningFabrikant` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `KoeretoejOplysningFrikoert` tinyint(1) DEFAULT NULL COMMENT 'KoeretoejOplysningFrikoert.xsd',
  `KoeretoejOplysningFredetForPladeInddragelse` tinyint(1) DEFAULT NULL COMMENT 'KoeretoejOplysningFredetForPladeInddragelse.xsd',
  `KoeretoejOplysningVejvenligLuftaffjedring` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejOplysningDanskGodkendelseNummer` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejOplysningAargang` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejOplysningIbrugtagningDato` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejOplysningTrafikskade` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejOplysningVeteranKoeretoejOriginal` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejOplysningEffektivitetforholdRelevant` tinyint(1) DEFAULT NULL COMMENT 'JaNejType.xsd',
  `KoeretoejOplysningEffektivitetforholdM3` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejOplysningEffektivitetforholdTon` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejOplysningVolumenorientering` tinyint(1) DEFAULT NULL COMMENT 'JaNejType.xsd',
  `KoeretoejOplysningSovekabine` tinyint(1) DEFAULT NULL COMMENT 'JaNejType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejoplysning_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejprisoplysninger`;
CREATE TABLE IF NOT EXISTS `koeretoejprisoplysninger` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `PrisOplysningerStandardPris` decimal(13,2) DEFAULT NULL COMMENT 'BeloebType.xsd',
  `PrisOplysningerIndkoebsPris` decimal(13,2) DEFAULT NULL COMMENT 'BeloebType.xsd',
  `PrisOplysningerMindsteBeskatningspris` decimal(13,2) DEFAULT NULL COMMENT 'BeloebType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejprisoplysninger_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejregistrering`;
CREATE TABLE IF NOT EXISTS `koeretoejregistrering` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejRegistreringNummer` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejRegistreringStatus` varchar(100) DEFAULT NULL COMMENT 'KoeretoejRegistreringStatusType.xsd',
  `KoeretoejRegistreringStatusDato` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `KoeretoejRegistreringStatusAarsag` varchar(40) DEFAULT NULL COMMENT 'KoeretoejRegistreringStatusAArsagType.xsd',
  `KoeretoejRegistreringKontrolTal` char(1) DEFAULT NULL COMMENT 'KodeEtCifferStartNulType.xsd',
  `KoeretoejRegistreringGyldigFra` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `KoeretoejRegistreringGyldigTil` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `KoeretoejRegistreringGrundlagIdent` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejRegistreringSenesteHaendelse` varchar(50) DEFAULT NULL COMMENT 'KoeretoejRegistreringHaendelserType.xsd',
  `KoeretoejRegistreringTilknyttetLeasingForhold` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejregistrering_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejregistreringgrundlag`;
CREATE TABLE IF NOT EXISTS `koeretoejregistreringgrundlag` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejRegistreringGrundlagIdent` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejRegistreringGrundlagStatus` varchar(100) DEFAULT NULL COMMENT 'KoeretoejRegistreringGrundlagStatusType.xsd',
  `KoeretoejRegistreringGrundlagStatusDato` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `KoeretoejRegistreringGrundlagGyldigFra` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `KoeretoejRegistreringGrundlagGyldigTil` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `KoeretoejRegistreringGrundlagKode` varchar(10) DEFAULT NULL COMMENT 'KodeType.xsd',
  `KoeretoejRegistreringGrundlagTilknyttetFasteProeveskilte` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejRegistreringGrundlagPeriodiskSyn` varchar(20) DEFAULT NULL COMMENT 'PeriodiskSynTypeType.xsd',
  `KoeretoejRegistreringGrundlagPeriodiskSynGyldigTil` date DEFAULT NULL COMMENT 'DatoType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejregistreringgrundlag_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejregistreringgrundlaganvendelse`;
CREATE TABLE IF NOT EXISTS `koeretoejregistreringgrundlaganvendelse` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejAnvendelseNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejAnvendelseNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejAnvendelseNummer`),
  CONSTRAINT `koeretoejregistreringgrundlaganvendelse_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejregistreringgrundlagart`;
CREATE TABLE IF NOT EXISTS `koeretoejregistreringgrundlagart` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejArtNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejArtNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejArtKraeverForsikring` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejArtBeskrivelse` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `KoeretoejArtGyldigFra` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejArtGyldigTil` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejArtStatus` varchar(40) DEFAULT NULL COMMENT 'ForretningParameterVaerdiStatusType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejArtNummer`),
  CONSTRAINT `koeretoejregistreringgrundlagart_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejregistreringgrundlaggenerelidentifikator`;
CREATE TABLE IF NOT EXISTS `koeretoejregistreringgrundlaggenerelidentifikator` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `RegistreringNummerNummer` varchar(10) DEFAULT NULL COMMENT 'RegistreringNummerType.xsd',
  `KoeretoejOplysningStelNummer` varchar(20) DEFAULT NULL COMMENT 'StelNummerType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejregistreringgrundlaggenerelidentifikator_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejregistreringnummer`;
CREATE TABLE IF NOT EXISTS `koeretoejregistreringnummer` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `RegistreringNummerIdent` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `RegistreringNummerType` varchar(40) DEFAULT NULL COMMENT 'RegistreringNummerTypeType.xsd',
  `RegistreringNummerStatus` varchar(40) DEFAULT NULL COMMENT 'RegistreringNummerStatusType.xsd',
  `RegistreringNummerStatusDato` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `RegistreringNummerKvadratiskIndhold1` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `RegistreringNummerKvadratiskIndhold2` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `RegistreringNummerAflangIndhold` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `RegistreringNummerUdloebDato` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `RegistreringNummerFigurantPlade` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `RegistreringNummerGraensepladeDkDato` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `RegistreringNummerNummer` varchar(10) DEFAULT NULL COMMENT 'RegistreringNummerType.xsd',
  `RegistreringNummerRettighedType` varchar(40) DEFAULT NULL COMMENT 'RegistreringNummerRettighedTypeType.xsd',
  `RegistreringNummerRettighedGyldigFra` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `RegistreringNummerRettighedGyldigTil` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `RegistreringNummerRettighedNummer` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `RegistreringNummerRettighedSidstAdviseretDato` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `RegistreringNummerRettighedKoerselFormaal` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `RegistreringNummerRettighedAntalFerieDageTilbage` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejregistreringnummer_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejregistreringnummerrettighed`;
CREATE TABLE IF NOT EXISTS `koeretoejregistreringnummerrettighed` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `RegistreringNummerRettighedType` varchar(40) DEFAULT NULL COMMENT 'RegistreringNummerRettighedTypeType.xsd',
  `RegistreringNummerRettighedGyldigFra` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `RegistreringNummerRettighedGyldigTil` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `RegistreringNummerRettighedNummer` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `RegistreringNummerRettighedSidstAdviseretDato` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `RegistreringNummerRettighedKoerselFormaal` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `RegistreringNummerRettighedAntalFerieDageTilbage` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejregistreringnummerrettighed_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejsupplerendekarrosseritype`;
CREATE TABLE IF NOT EXISTS `koeretoejsupplerendekarrosseritype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `SupplerendeKarrosseriTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `SupplerendeKarrosseriTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`SupplerendeKarrosseriTypeNummer`),
  CONSTRAINT `koeretoejsupplerendekarrosseritype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejsynresultat`;
CREATE TABLE IF NOT EXISTS `koeretoejsynresultat` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejMotorKilometerstand` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `SynResultatNummer` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `SynResultatSynStatus` varchar(50) DEFAULT NULL COMMENT 'SynStatusType.xsd',
  `SynResultatSynStatusDato` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `SynResultatSynsDato` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `SynResultatSynsResultat` varchar(60) DEFAULT NULL COMMENT 'SynResultatType.xsd',
  `SynResultatSynsType` varchar(50) DEFAULT NULL COMMENT 'SynTypeType.xsd',
  `SynResultatOmsynMoedeDato` date DEFAULT NULL COMMENT 'DatoType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejsynresultat_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtilladelse`;
CREATE TABLE IF NOT EXISTS `koeretoejtilladelse` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `TilladelseNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `TilladelseGyldigFra` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `TilladelseGyldigTil` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `TilladelseKommentar` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `TilladelseKunGodkendtForRegistreretEjer` varchar(20) DEFAULT NULL COMMENT 'JuridiskEnhedIDType.xsd',
  `TilladelseKombinationKoeretoejIdent` bigint(16) DEFAULT NULL COMMENT 'KidType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`TilladelseNummer`),
  CONSTRAINT `koeretoejtilladelse_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtilladelsetype`;
CREATE TABLE IF NOT EXISTS `koeretoejtilladelsetype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `TilladelseTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `TilladelseTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `TilladelseTypeErPeriodeBegraenset` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `TilladelseTypePeriodeLaengde` int(6) DEFAULT NULL COMMENT 'PeriodeLaengdeType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`TilladelseTypeNummer`),
  CONSTRAINT `koeretoejtilladelsetype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtilladelsetypefasttilkobling`;
CREATE TABLE IF NOT EXISTS `koeretoejtilladelsetypefasttilkobling` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejFastTilkoblingIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejOplysningStelNummer` varchar(20) DEFAULT NULL COMMENT 'StelNummerType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejFastTilkoblingIdent`),
  CONSTRAINT `koeretoejtilladelsetypefasttilkobling_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtilladelsetypekungodkendtforjuridiskenhed`;
CREATE TABLE IF NOT EXISTS `koeretoejtilladelsetypekungodkendtforjuridiskenhed` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `PersonCPRNummer` varchar(10) DEFAULT NULL COMMENT 'CPRNummerType.xsd',
  `VirksomhedSENummer` char(8) DEFAULT NULL COMMENT 'SENummerType.xsd',
  `VirksomhedCVRNummer` char(8) DEFAULT NULL COMMENT 'CVRNummerType.xsd',
  `ProduktionEnhedNummer` bigint(10) DEFAULT NULL COMMENT 'ProduktionEnhedNummerType.xsd',
  `AlternativKontaktID` int(9) DEFAULT NULL COMMENT 'TalHel9Type.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejtilladelsetypekungodkendtforjuridiskenhed_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtilladelsetypevariabelkombination`;
CREATE TABLE IF NOT EXISTS `koeretoejtilladelsetypevariabelkombination` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejVariabelKombinationIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `RegistreringNummerNummer` varchar(10) DEFAULT NULL COMMENT 'RegistreringNummerType.xsd',
  `KoeretoejOplysningStelNummer` varchar(20) DEFAULT NULL COMMENT 'StelNummerType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejVariabelKombinationIdent`),
  CONSTRAINT `koeretoejtilladelsetypevariabelkombination_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattest`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattest` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `TypeAttestGyldigFra` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `TypeAttestGyldigTil` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `TypeAttestTypeGodkendelseNummer` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `TypeAttestTypeAnmeldelseNummer` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejtypeattest_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattestkoeretoejart`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattestkoeretoejart` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejArtNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejArtNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejArtKraeverForsikring` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejArtBeskrivelse` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `KoeretoejArtGyldigFra` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejArtGyldigTil` date DEFAULT NULL COMMENT 'DatoType.xsd',
  `KoeretoejArtStatus` varchar(40) DEFAULT NULL COMMENT 'ForretningParameterVaerdiStatusType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejArtNummer`),
  CONSTRAINT `koeretoejtypeattestkoeretoejart_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattestkoeretoejmaerke`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattestkoeretoejmaerke` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejMaerkeTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejMaerkeTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejMaerkeTypeNummer`),
  CONSTRAINT `koeretoejtypeattestkoeretoejmaerke_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattestkoeretoejmodel`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattestkoeretoejmodel` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejModelTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejModelTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejModelTypeNummer`),
  CONSTRAINT `koeretoejtypeattestkoeretoejmodel_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattestkoeretoejtype`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattestkoeretoejtype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejTypeTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejTypeTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejTypeTypeNummer`),
  CONSTRAINT `koeretoejtypeattestkoeretoejtype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattestkoeretoejvariant`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattestkoeretoejvariant` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejVariantTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejVariantTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejVariantTypeNummer`),
  CONSTRAINT `koeretoejtypeattestkoeretoejvariant_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattesttilladelse`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattesttilladelse` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `TilladelseTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `TilladelseTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `TilladelseTypeErPeriodeBegraenset` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `TilladelseTypePeriodeLaengde` int(6) DEFAULT NULL COMMENT 'PeriodeLaengdeType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`TilladelseTypeNummer`),
  CONSTRAINT `koeretoejtypeattesttilladelse_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattestvariant`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattestvariant` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `TypeAttestVariantNummer` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantSiddepladserMinimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantSiddepladserMaksimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantEgenVaegt` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantKoereklarVaegtMaksimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantTekniskTotalVaegt` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantTotalVaegt` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantStoersteAkselTryk` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantTilkoblingsvaegtMedBremser` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantTilkoblingsvaegtUdenBremser` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantStatus` varchar(7) DEFAULT NULL COMMENT 'TypeAttestVariantStatusType.xsd',
  `TypeAttestVariantStatusDatoTid` datetime DEFAULT NULL COMMENT 'DatoTidType.xsd',
  `TypeAttestVariantStaapladserMinimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantStaapladserMaksimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantPassagerAntal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantAkselAntal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantFaelgDaek` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `TypeAttestVariantMaksimumHastighed` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantStelNummerAnbringelse` varchar(500) DEFAULT NULL COMMENT 'TekstLangType.xsd',
  `TypeAttestVariantVVaerdiLuft` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantVVaerdiMekanisk` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantTraekkendeAksler` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `TypeAttestVariantAntalGear` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantAntalDoere` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantCO2Udslip` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantRoegtaethed` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantRoegtaethedOmdrejningstal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantPartikelFilter` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantCylinderAntal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantMaerkning` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `TypeAttestVariantStandStoej` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantStandStoejOmdrejningstal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantKoerselStoej` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantKoereklarVaegtMinimum` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantEgnetTilTaxi` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `TypeAttestVariantPartikler` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantKmPerLiter` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantStoersteEffekt` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantInnovativTeknik` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `TypeAttestVariantInnovativTeknikAntal` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantNCAPTest` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `TypeAttestVariantSkammelBelastning` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantSkatteAkselAntal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantSkatteAkselTryk` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantSaettevognTilladtAkselTryk` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantVogntogVaegt` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantAkselAfstand` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantSporviddenForrest` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantSporviddenBagest` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `TypeAttestVariantSlagVolumen` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `TypeAttestVariantElektriskForbrug` float DEFAULT NULL COMMENT 'TalFlydendeType.xsd',
  `KoeretoejVariantTypeNummer` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejTypeTypeNummer` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`),
  CONSTRAINT `koeretoejtypeattestvariant_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattestvariantdrivkraft`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattestvariantdrivkraft` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `DrivkraftTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `DrivkraftTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`DrivkraftTypeNummer`),
  CONSTRAINT `koeretoejtypeattestvariantdrivkraft_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattestvariantkarrosseri`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattestvariantkarrosseri` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KarrosseriTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KarrosseriTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KarrosseriTypeNummer`),
  CONSTRAINT `koeretoejtypeattestvariantkarrosseri_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypeattestvariantnorm`;
CREATE TABLE IF NOT EXISTS `koeretoejtypeattestvariantnorm` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `NormTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `NormTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`NormTypeNummer`),
  CONSTRAINT `koeretoejtypeattestvariantnorm_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejtypetype`;
CREATE TABLE IF NOT EXISTS `koeretoejtypetype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejTypeTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejTypeTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejTypeTypeNummer`),
  CONSTRAINT `koeretoejtypetype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejudstyr`;
CREATE TABLE IF NOT EXISTS `koeretoejudstyr` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejUdstyrAntal` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejUdstyrTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejUdstyrTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  `KoeretoejUdstyrTypeVisesVedSyn` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejUdstyrTypeVisesVedForespoergsel` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejUdstyrTypeVisesVedStandardOprettelse` tinyint(1) DEFAULT NULL COMMENT 'BolskType.xsd',
  `KoeretoejUdstyrTypeStandardAntal` bigint(18) DEFAULT NULL COMMENT 'TalHelType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejUdstyrAntal`,`KoeretoejUdstyrTypeNummer`),
  CONSTRAINT `koeretoejudstyr_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejundergruppe`;
CREATE TABLE IF NOT EXISTS `koeretoejundergruppe` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejUndergruppeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejUndergruppeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejUndergruppeNummer`),
  CONSTRAINT `koeretoejundergruppe_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

DROP TABLE IF EXISTS `koeretoejvarianttype`;
CREATE TABLE IF NOT EXISTS `koeretoejvarianttype` (
  `KoeretoejIdent` bigint(16) NOT NULL COMMENT 'KidType.xsd',
  `KoeretoejVariantTypeNummer` bigint(18) NOT NULL COMMENT 'TalHelType.xsd',
  `KoeretoejVariantTypeNavn` varchar(100) DEFAULT NULL COMMENT 'TekstKortType.xsd',
  UNIQUE KEY `KoeretoejIdent` (`KoeretoejIdent`,`KoeretoejVariantTypeNummer`),
  CONSTRAINT `koeretoejvarianttype_ibfk_1` FOREIGN KEY (`KoeretoejIdent`) REFERENCES `koeretoej` (`KoeretoejIdent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_danish_ci;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
