package main

import (
	"archive/zip"
	"database/sql"
	"context"
	// "encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	// "io/ioutil"
	"sync"
	"sync/atomic"
	"strings"
	"time"
	"runtime"
	"log"

	_ "github.com/go-sql-driver/mysql"
)





// KoeretoejAnvendelseStrukturType.xsd
type KoeretoejAnvendelseSamlingStruct struct {
	KoeretoejAnvendelse 							  			  []KoeretoejAnvendelseStruct `xml:"KoeretoejAnvendelseStruktur"`
}

// KoeretoejAnvendelse.xsd
type KoeretoejAnvendelseStruct struct {
	KoeretoejAnvendelseNavn                           			  string `xml:"KoeretoejAnvendelseNavn"`
	KoeretoejAnvendelseNummer                         			  string `xml:"KoeretoejAnvendelseNummer"`
	KoeretoejAnvendelseBeskrivelse                    			  string `xml:"KoeretoejAnvendelseBeskrivelse"`
	KoeretoejAnvendelseGyldigFra                      			  string `xml:"KoeretoejAnvendelseGyldigFra"`
	KoeretoejAnvendelseGyldigTil                      			  string `xml:"KoeretoejAnvendelseGyldigTil"`
	KoeretoejAnvendelseStatus                         			  string `xml:"KoeretoejAnvendelseStatus"`
}




// DrivmiddelStrukturType.xsd
type KoeretoejDrivmiddelSamlingStruct struct {
	DrivmiddelStruktur 								  			  []DrivmiddelStrukturStruct `xml:"DrivmiddelStruktur"`
}

// DrivmiddelStrukturType.xsd
type DrivmiddelStrukturStruct struct {
	KoeretoejMotorBraendselscelle                     			  string `xml:"KoeretoejMotorBraendselscelle"`
	KoeretoejMotorPlugInHybrid                        			  string `xml:"KoeretoejMotorPlugInHybrid"`
	KoeretoejMotorDrivmiddelPrimaer                   			  string `xml:"KoeretoejMotorDrivmiddelPrimaer"`

	// MaaleNormStrukturType.xsd
	KoeretoejMotorMaaleNormTypeNavn                   			  string `xml:"MaaleNormStruktur>KoeretoejMotorMaaleNormTypeNavn"`
	KoeretoejMotorMaaleNormTypeNummer 		  		  			  string `xml:"MaaleNormStruktur>KoeretoejMotorMaaleNormTypeNummer"`

	// DrivkraftType.xsd
	DrivkraftTypeNavn                   		  	  			  string `xml:"DrivkraftTypeStruktur>DrivkraftTypeNavn"`
	DrivkraftTypeNummer 		  				      			  string `xml:"DrivkraftTypeStruktur>DrivkraftTypeNummer"`

	// KoeretoejBraendstofStrukturType.xsd
	KoeretoejMotorKmPerLiter 		  				  			  string `xml:"KoeretoejBraendstofStruktur>KoeretoejMotorKmPerLiter"`
	KoeretoejMotorKMPerLiterPreCalc 		  		  			  string `xml:"KoeretoejBraendstofStruktur>KoeretoejMotorKMPerLiterPreCalc"`
	KoeretoejMotorBraendstofforbrugMaalt 		  	  			  string `xml:"KoeretoejBraendstofStruktur>KoeretoejMotorBraendstofforbrugMaalt"`
	KoeretoejMotorGasforbrug 		  		          			  string `xml:"KoeretoejBraendstofStruktur>KoeretoejMotorGasforbrug"`
	KoeretoejMotorCO2UdslipBeregnet 		  		  			  string `xml:"KoeretoejBraendstofStruktur>KoeretoejMotorCO2UdslipBeregnet"`
	KoeretoejMiljoeOplysningCO2Udslip 		  		  			  string `xml:"KoeretoejBraendstofStruktur>KoeretoejMiljoeOplysningCO2Udslip"`

	// KoeretoejElforbrugStrukturType.xsd
	KoeretoejMotorElektriskForbrug                    			  string `xml:"KoeretoejElforbrugStruktur>KoeretoejMotorElektriskForbrug"`
	KoeretoejMotorElektriskForbrugMaalt               			  string `xml:"KoeretoejElforbrugStruktur>KoeretoejMotorElektriskForbrugMaalt"`
	KoeretoejMotorElektriskRaekkevidde                			  string `xml:"KoeretoejElforbrugStruktur>KoeretoejMotorElektriskRaekkevidde"`
	KoeretoejMotorBatterikapacitet                    			  string `xml:"KoeretoejElforbrugStruktur>KoeretoejMotorBatterikapacitet"`
}




// DispensationTypeStrukturType.xsd
type DispensationTypeSamlingStruct struct {
	DispensationTypeStruktur 						  			  []DispensationTypeStrukturStruct `xml:"DispensationTypeStruktur"`
}

// DispensationType.xsd
type DispensationTypeStrukturStruct struct {
	DispensationTypeNummer                            			  string `xml:"DispensationTypeNummer"`
	DispensationTypeNavn                              			  string `xml:"DispensationTypeNavn"`
	KoeretoejDispensationTypeKommentar                			  string `xml:"KoeretoejDispensationTypeKommentar"`
}




// KoeretoejSupplerendeKarrosseriTypeStrukturType.xsd
type KoeretoejSupplerendeKarrosseriSamlingStruct struct {
	KoeretoejSupplerendeKarrosseriTypeStruktur 		  			  []KoeretoejSupplerendeKarrosseriTypeStrukturStruct `xml:"KoeretoejSupplerendeKarrosseriTypeStruktur"`
}

// SupplerendeKarrosseriType.xsd
type KoeretoejSupplerendeKarrosseriTypeStrukturStruct struct {
	SupplerendeKarrosseriTypeNavn   				  			  string `xml:"SupplerendeKarrosseriTypeNavn"`
	SupplerendeKarrosseriTypeNummer 				  			  string `xml:"SupplerendeKarrosseriTypeNummer"`
}




// KoeretoejUdstyrStrukturType.xsd
type KoeretoejUdstyrSamlingStruct struct {
	KoeretoejUdstyrStruktur 						  			  []KoeretoejUdstyrStrukturStruct `xml:"KoeretoejUdstyrStruktur"`
}

// KoeretoejUdstyrStrukturType.xsd
type KoeretoejUdstyrStrukturStruct struct {
	KoeretoejUdstyrAntal   							  			  string `xml:"KoeretoejUdstyrAntal"`

	// KoeretoejUdstyrType.xsd
	KoeretoejUdstyrTypeNavn 						  			  string `xml:"KoeretoejUdstyrTypeStruktur>KoeretoejUdstyrTypeNavn"`
	KoeretoejUdstyrTypeNummer 						  			  string `xml:"KoeretoejUdstyrTypeStruktur>KoeretoejUdstyrTypeNummer"`
	KoeretoejUdstyrTypeStandardAntal 				  			  string `xml:"KoeretoejUdstyrTypeStruktur>KoeretoejUdstyrTypeStandardAntal"`
	KoeretoejUdstyrTypeVisesVedForespoergsel 		  			  string `xml:"KoeretoejUdstyrTypeStruktur>KoeretoejUdstyrTypeVisesVedForespoergsel"`
	KoeretoejUdstyrTypeVisesVedStandardOprettelse 	  			  string `xml:"KoeretoejUdstyrTypeStruktur>KoeretoejUdstyrTypeVisesVedStandardOprettelse"`
	KoeretoejUdstyrTypeVisesVedSyn 					  			  string `xml:"KoeretoejUdstyrTypeStruktur>KoeretoejUdstyrTypeVisesVedSyn"`
}




// KoeretoejBlokeringAArsagTypeStrukturType.xsd
type KoeretoejBlokeringAarsagListeStruct struct {
	KoeretoejBlokeringAarsag 						  			  []KoeretoejBlokeringStruct `xml:"KoeretoejBlokeringAarsag"`
}

// KoeretoejBlokeringAArsagType.xsd
type KoeretoejBlokeringStruct struct {
	KoeretoejBlokeringAarsagTypeNavn   				  			  string `xml:"KoeretoejBlokeringAarsagTypeNavn"`
	KoeretoejBlokeringAarsagTypeNummer 				  			  string `xml:"KoeretoejBlokeringAarsagTypeNummer"`
}




// TilladelseSamlingStrukturType.xsd
type TilladelseStruct struct {
	TilladelseNummer                                  			  string `xml:"TilladelseNummer"`

	TilladelseStruktur 								  			  []TilladelseStrukturStruct `xml:"TilladelseStruktur"`
}

// TilladelseStrukturType.xsd
type TilladelseStrukturStruct struct {
	// Tilladelse.xsd
	TilladelseGyldigFra   							  			  string `xml:"TilladelseGyldigFra"`
	TilladelseGyldigTil                               			  string `xml:"TilladelseGyldigTil"`
	TilladelseKommentar   							  			  string `xml:"TilladelseKommentar"`
	TilladelseNummer                                  			  string `xml:"TilladelseNummer"`
	TilladelseKunGodkendtForRegistreretEjer           			  string `xml:"TilladelseKunGodkendtForRegistreretEjer"`
	TilladelseKombinationKoeretoejIdent               			  string `xml:"TilladelseKombinationKoeretoejIdent"`

	// TilladelseType.xsd
	TilladelseTypeNavn 								  			  string `xml:"TilladelseTypeStruktur>TilladelseTypeNavn"`
	TilladelseTypeNummer 							  			  string `xml:"TilladelseTypeStruktur>TilladelseTypeNummer"`
	TilladelseTypeErPeriodeBegraenset                 			  string `xml:"TilladelseTypeStruktur>TilladelseTypeErPeriodeBegraenset"`
	TilladelseTypePeriodeLaengde                      			  string `xml:"TilladelseTypeStruktur>TilladelseTypePeriodeLaengde"`

	// KoeretoejGenerelIdentifikatorStrukturType.xsd
	VariabelKombinationKoeretoejIdent                 			  string `xml:"TilladelseTypeDetaljeValg>VariabelKombination>KoeretoejGenerelIdentifikatorStruktur>KoeretoejGenerelIdentifikatorValg>KoeretoejIdent"`
	VariabelKombinationRegistreringNummerNummer       			  string `xml:"TilladelseTypeDetaljeValg>VariabelKombination>KoeretoejGenerelIdentifikatorStruktur>KoeretoejGenerelIdentifikatorValg>RegistreringNummerNummer"`
	VariabelKombinationKoeretoejOplysningStelNummer   			  string `xml:"TilladelseTypeDetaljeValg>VariabelKombination>KoeretoejGenerelIdentifikatorStruktur>KoeretoejGenerelIdentifikatorValg>KoeretoejOplysningStelNummer"`

	// TilladelseStrukturType.xsd
	FastTilkoblingKoeretoejIdent                      			  string `xml:"TilladelseTypeDetaljeValg>FastTilkobling>KoeretoejIdent"`
	FastTilkoblingKoeretoejOplysningStelNummer        			  string `xml:"TilladelseTypeDetaljeValg>FastTilkobling>KoeretoejOplysningStelNummer"`

	/* FUTURE WORKS: JuridiskEnhedStrukturType.xsd
	string `xml:"TilladelseTypeDetaljeValg>KunGodkendtForJuridiskEnhed>JuridiskEnhedIdentifikatorStruktur>JuridiskEnhedValg>"`
	*/
}




// ESStatistikListeModtag_IType.xsd
type Statistik struct {
	KoeretoejIdent                            		  			  string `xml:"KoeretoejIdent"`

	// KoeretoejArt.xsd
	KoeretoejArtNavn                          		  			  string `xml:"KoeretoejArtNavn"`
	KoeretoejArtNummer                        		  			  string `xml:"KoeretoejArtNummer"`
	KoeretoejArtKraeverForsikring                     			  string `xml:"KoeretoejArtKraeverForsikring"`
	KoeretoejArtBeskrivelse                           			  string `xml:"KoeretoejArtBeskrivelse"`
	KoeretoejArtGyldigFra                             			  string `xml:"KoeretoejArtGyldigFra"`
	KoeretoejArtGyldigTil                             			  string `xml:"KoeretoejArtGyldigTil"`
	KoeretoejArtStatus                                			  string `xml:"KoeretoejArtStatus"`

	// KoeretoejAnvendelse.xsd
	KoeretoejAnvendelseNavn                           			  string `xml:"KoeretoejAnvendelseStruktur>KoeretoejAnvendelseNavn"`
	KoeretoejAnvendelseNummer                         			  string `xml:"KoeretoejAnvendelseStruktur>KoeretoejAnvendelseNummer"`
	KoeretoejAnvendelseBeskrivelse                    			  string `xml:"KoeretoejAnvendelseStruktur>KoeretoejAnvendelseBeskrivelse"`
	KoeretoejAnvendelseGyldigFra                      			  string `xml:"KoeretoejAnvendelseStruktur>KoeretoejAnvendelseGyldigFra"`
	KoeretoejAnvendelseGyldigTil                      			  string `xml:"KoeretoejAnvendelseStruktur>KoeretoejAnvendelseGyldigTil"`
	KoeretoejAnvendelseStatus                         			  string `xml:"KoeretoejAnvendelseStruktur>KoeretoejAnvendelseStatus"`

	// KoeretoejAnvendelseSamlingStrukturType.xsd
	KoeretoejAnvendelseSamling 						  			  KoeretoejAnvendelseSamlingStruct `xml:"KoeretoejAnvendelseSamlingStruktur>KoeretoejAnvendelseSamling"`

	// Leasing.xsd
	LeasingMaaneder                                   			  string `xml:"LeasingMaaneder"`
	LeasingNummer                                     			  string `xml:"LeasingNummer"`
	LeasingGyldigFra                  				  			  string `xml:"LeasingGyldigFra"`
	LeasingGyldigTil                  				  			  string `xml:"LeasingGyldigTil"`
	LeasingReelOphoerDato                             			  string `xml:"LeasingReelOphoerDato"`
	LeasingKode                                       			  string `xml:"LeasingKode"`
	LeasingStatus                                     			  string `xml:"LeasingStatus"`
	LeasingBemaerkning                                			  string `xml:"LeasingBemaerkning"`
	LeasingAendringType                               			  string `xml:"LeasingAendringType"`
	LeasingSidstAendret                               			  string `xml:"LeasingSidstAendret"`

	// RegistreringNummer.xsd
	RegistreringNummerIdent                           			  string `xml:"RegistreringNummerIdent"`
	RegistreringNummerAflangIndhold                   			  string `xml:"RegistreringNummerAflangIndhold"`
	RegistreringNummerGraensepladeDkDato              			  string `xml:"RegistreringNummerGraensepladeDkDato"`
	RegistreringNummerKvadratiskIndhold1              			  string `xml:"RegistreringNummerKvadratiskIndhold1"`
	RegistreringNummerKvadratiskIndhold2              			  string `xml:"RegistreringNummerKvadratiskIndhold2"`
	RegistreringNummerNummer                  		  			  string `xml:"RegistreringNummerNummer"`
	RegistreringNummerStatus                          			  string `xml:"RegistreringNummerStatus"`
	RegistreringNummerStatusDato                      			  string `xml:"RegistreringNummerStatusDato"`
	RegistreringNummerType                            			  string `xml:"RegistreringNummerType"`
	RegistreringNummerUdloebDato                  	  			  string `xml:"RegistreringNummerUdloebDato"`
	RegistreringNummerFigurantPlade                   			  string `xml:"RegistreringNummerFigurantPlade"`

	// RegistreringNummerRettighed.xsd
	RegistreringNummerRettighedGyldigFra              			  string `xml:"RegistreringNummerRettighedGyldigFra"`
	RegistreringNummerRettighedGyldigTil              			  string `xml:"RegistreringNummerRettighedGyldigTil"`
	RegistreringNummerRettighedNummer                 			  string `xml:"RegistreringNummerRettighedNummer"`
	RegistreringNummerRettighedSidstAdviseretDato     			  string `xml:"RegistreringNummerRettighedSidstAdviseretDato"`
	RegistreringNummerRettighedType                   			  string `xml:"RegistreringNummerRettighedType"`
	RegistreringNummerRettighedKoerselFormaal         			  string `xml:"RegistreringNummerRettighedKoerselFormaal"`
	RegistreringNummerRettighedAntalFerieDageTilbage  			  string `xml:"RegistreringNummerRettighedAntalFerieDageTilbage"`

	// KoeretoejOplysningStrukturType.xsd (KoeretoejOplysningFrikoert & KoeretoejOplysningFredetForPladeInddragelse)
	// KoeretoejOplysning.xsd
	KoeretoejOplysningOprettetUdFra                	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningOprettetUdFra"`
	KoeretoejOplysningStatus                       	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningStatus"`
	KoeretoejOplysningStatusDato                   	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningStatusDato"`
	 KoeretoejOplysningFoersteRegistreringDato      	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningFoersteRegistreringDato"`
	KoeretoejOplysningStelNummer                   	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningStelNummer"`
	KoeretoejOplysningStelNummerAnbringelse        	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningStelNummerAnbringelse"`
	KoeretoejOplysningModelAar                     	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningModelAar"`
	KoeretoejOplysningTotalVaegt                   	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTotalVaegt"`
	KoeretoejOplysningEgenVaegt                    	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningEgenVaegt"`
	KoeretoejOplysningKoereklarVaegtMinimum        	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningKoereklarVaegtMinimum"`
	KoeretoejOplysningKoereklarVaegtMaksimum       	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningKoereklarVaegtMaksimum"`
	KoeretoejOplysningTekniskTotalVaegt            	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTekniskTotalVaegt"`
	KoeretoejOplysningVogntogVaegt                 	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningVogntogVaegt"`
	KoeretoejOplysningAkselAntal                   	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningAkselAntal"`
	KoeretoejOplysningStoersteAkselTryk            	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningStoersteAkselTryk"`
	KoeretoejOplysningSkatteAkselAntal             	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningSkatteAkselAntal"`
	KoeretoejOplysningSkatteAkselTryk              	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningSkatteAkselTryk"`
	KoeretoejOplysningPassagerAntal                	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningPassagerAntal"`
	KoeretoejOplysningSiddepladserMinimum          	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningSiddepladserMinimum"`
	KoeretoejOplysningSiddepladserMaksimum         	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningSiddepladserMaksimum"`
	KoeretoejOplysningStaapladserMinimum           	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningStaapladserMinimum"`
	KoeretoejOplysningStaapladserMaksimum          	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningStaapladserMaksimum"`
	KoeretoejOplysningTilkoblingMulighed           	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTilkoblingMulighed"`
	KoeretoejOplysningTilkoblingsvaegtUdenBremser  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTilkoblingsvaegtUdenBremser"`
	KoeretoejOplysningTilkoblingsvaegtMedBremser   	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTilkoblingsvaegtMedBremser"`
	KoeretoejOplysningPaahaengVognTotalVaegt       	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningPaahaengVognTotalVaegt"`
	KoeretoejOplysningSkammelBelastning            	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningSkammelBelastning"`
	KoeretoejOplysningSaettevognTilladtAkselTryk   	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningSaettevognTilladtAkselTryk"`
	KoeretoejOplysningMaksimumHastighed            	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningMaksimumHastighed"`
	KoeretoejOplysningFaelgDaek                    	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningFaelgDaek"`
	KoeretoejOplysningTilkobletSidevognStelnr      	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTilkobletSidevognStelnr"`
	KoeretoejOplysningNCAPTest                     	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningNCAPTest"`
	KoeretoejOplysningVVaerdiLuft                  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningVVaerdiLuft"`
	KoeretoejOplysningVVaerdiMekanisk              	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningVVaerdiMekanisk"`
	KoeretoejOplysningOevrigtUdstyr                	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningOevrigtUdstyr"`
	KoeretoejOplysningKoeretoejstand               	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningKoeretoejstand"`
	KoeretoejOplysning30PctVarevogn                	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysning30PctVarevogn"`
	KoeretoejOplysningBlokvognAkselType				  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognAkselType"`
	KoeretoejOplysningBlokvognHovedboltTryk			  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognHovedboltTryk"`
	KoeretoejOplysningBlokvognSkammelTryk			  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognSkammelTryk"`
	KoeretoejOplysningBlokvognSamletAkselTryk		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognSamletAkselTryk"`
	KoeretoejOplysningBlokvognMaxVogntog			  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognMaxVogntog"`
	KoeretoejOplysningBlokvognBreddeFra				  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognBreddeFra"`
	KoeretoejOplysningBlokvognKoblingshoejdeFra		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognKoblingshoejdeFra"`
	KoeretoejOplysningBlokvognKoblingslaengdeFra	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognKoblingslaengdeFra"`
	KoeretoejOplysningBlokvognSammenkoblingType		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognSammenkoblingType"`
	KoeretoejOplysningBlokvognTilladeligHastighed	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognTilladeligHastighed"`
	KoeretoejOplysningBlokvognBreddeTil				  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognBreddeTil"`
	KoeretoejOplysningBlokvognKoblingshoejdeTil		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognKoblingshoejdeTil"`
	KoeretoejOplysningBlokvognKoblingslaengdeTil	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningBlokvognKoblingslaengdeTil"`
	KoeretoejOplysningTraekkendeAksler             	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTraekkendeAksler"`
	KoeretoejOplysningEgnetTilTaxi                 	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningEgnetTilTaxi"`
	KoeretoejOplysningAkselAfstand                 	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningAkselAfstand"`
	KoeretoejOplysningSporviddenForrest            	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningSporviddenForrest"`
	KoeretoejOplysningSporviddenBagest             	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningSporviddenBagest"`
	KoeretoejOplysningTypeAnmeldelseNummer         	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTypeAnmeldelseNummer"`
	KoeretoejOplysningTypeGodkendelseNummer        	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTypeGodkendelseNummer"`
	KoeretoejOplysningEUVariant                    	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningEUVariant"`
	KoeretoejOplysningEUVersion                    	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningEUVersion"`
	KoeretoejOplysningKommentar                    	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningKommentar"`
	KoeretoejOplysningTypegodkendtKategori         	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTypegodkendtKategori"`
	KoeretoejOplysningAntalGear                    	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningAntalGear"`
	KoeretoejOplysningAntalDoere                   	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningAntalDoere"`
	KoeretoejOplysningFabrikant                    	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningFabrikant"`
	KoeretoejOplysningFrikoert                    	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningFrikoert"`
	KoeretoejOplysningFredetForPladeInddragelse       			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningFredetForPladeInddragelse"`
	KoeretoejOplysningVejvenligLuftaffjedring         			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningVejvenligLuftaffjedring"`
	KoeretoejOplysningDanskGodkendelseNummer          			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningDanskGodkendelseNummer"`
	KoeretoejOplysningAargang                    	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningAargang"`
	KoeretoejOplysningIbrugtagningDato             	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningIbrugtagningDato"`
	KoeretoejOplysningTrafikskade                  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningTrafikskade"`
	KoeretoejOplysningVeteranKoeretoejOriginal     	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningVeteranKoeretoejOriginal"`
	KoeretoejOplysningEffektivitetforholdRelevant     			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningEffektivitetforholdRelevant"`
	KoeretoejOplysningEffektivitetforholdM3           			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningEffektivitetforholdM3"`
	KoeretoejOplysningEffektivitetforholdTon          			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningEffektivitetforholdTon"`
	KoeretoejOplysningVolumenorientering              			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningVolumenorientering"`
	KoeretoejOplysningSovekabine                      			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejOplysningSovekabine"`

	// KoeretoejOplysningGrundStrukturType.xsd
	KoeretoejFastKombinationKoeretoejIdent            			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejFastKombination>KoeretoejIdent"`
	KoeretoejFastKombinationRegistreringNummerNummer  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejFastKombination>RegistreringNummerNummer"`
	KoeretoejFastKombinationRegistreringNummerIdent   			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejFastKombination>RegistreringNummerIdent"`

	// KoeretoejBetegnelseStrukturType.xsd
	KoeretoejMaerkeTypeNavn                   		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejBetegnelseStruktur>KoeretoejMaerkeTypeNavn"`
	KoeretoejMaerkeTypeNummer                 		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejBetegnelseStruktur>KoeretoejMaerkeTypeNummer"`
	KoeretoejModelTypeNavn                    		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejBetegnelseStruktur>Model>KoeretoejModelTypeNavn"`
	KoeretoejModelTypeNummer                  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejBetegnelseStruktur>Model>KoeretoejModelTypeNummer"`
	KoeretoejTypeTypeNavn                     		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejBetegnelseStruktur>Type>KoeretoejTypeTypeNavn"`
	KoeretoejTypeTypeNummer                   		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejBetegnelseStruktur>Type>KoeretoejTypeTypeNummer"`
	KoeretoejVariantTypeNavn                  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejBetegnelseStruktur>Variant>KoeretoejVariantTypeNavn"`
	KoeretoejVariantTypeNummer                		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejBetegnelseStruktur>Variant>KoeretoejVariantTypeNummer"`

	// FarveType.xsd
	FarveTypeNavn                  			  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejFarveStruktur>FarveTypeStruktur>FarveTypeNavn"`
	FarveTypeNummer                			  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejFarveStruktur>FarveTypeStruktur>FarveTypeNummer"`

	// KarrosseriType.xsd
	KarrosseriTypeNavn                   	  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KarrosseriTypeStruktur>KarrosseriTypeNavn"`
	KarrosseriTypeNummer                   	  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KarrosseriTypeStruktur>KarrosseriTypeNummer"`

	// KoeretoejSupplerendeKarrosseriSamlingStrukturType.xsd
	KoeretoejSupplerendeKarrosseriSamling 			  			  KoeretoejSupplerendeKarrosseriSamlingStruct `xml:"KoeretoejOplysningGrundStruktur>KoeretoejSupplerendeKarrosseriSamlingStruktur>KoeretoejSupplerendeKarrosseriSamling"`

	// NormType.xsd
	NormTypeNavn 		  		  	  				  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejNormStruktur>NormTypeStruktur>NormTypeNavn"`
	NormTypeNummer 	  								  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejNormStruktur>NormTypeStruktur>NormTypeNummer"`

	// KoeretoejMiljoeOplysning.xsd
	KoeretoejMiljoeOplysningCO2Udslip				  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningCO2Udslip"`
	KoeretoejMiljoeOplysningEmissionCO        		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningEmissionCO"`
	KoeretoejMiljoeOplysningEmissionHCPlusNOX 		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningEmissionHCPlusNOX"`
	KoeretoejMiljoeOplysningEmissionNOX       		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningEmissionNOX"`
	KoeretoejMiljoeOplysningPartikler 		  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningPartikler"`
	KoeretoejMiljoeOplysningPartikelFilter    		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningPartikelFilter"`
	KoeretoejMiljoeOplysningRoegtaethed       		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningRoegtaethed"`
	KoeretoejMiljoeOplysningRoegtaethedOmdrejningstal 			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningRoegtaethedOmdrejningstal"`
	KoeretoejMiljoeOplysningEftermonteretPartikelfilter 		  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningEftermonteretPartikelfilter"`
	KoeretoejMiljoeOplysningSpecifikCO2Emission       			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningSpecifikCO2Emission"`
	KoeretoejMiljoeOplysningNyttelastvaerdi        	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningNyttelastvaerdi"`
	KoeretoejMiljoeOplysningTungtNulEmissionKoeretoej 			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningTungtNulEmissionKoeretoej"`
	KoeretoejMiljoeOplysningCO2EmissionKlasseOmklassificeringDato string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningCO2EmissionKlasseOmklassificeringDato"`
	KoeretoejMiljoeOplysningCO2EmissionKlasseOprindelig 		  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>KoeretoejMiljoeOplysningCO2EmissionKlasseOprindelig"`

	// CO2EmissionKlasse.xsd
	CO2EmissionKlasseNavn				              			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>CO2EmissionKlasseStruktur>CO2EmissionKlasseNavn"`
	CO2EmissionKlasseNummer				              			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMiljoeOplysningStruktur>CO2EmissionKlasseStruktur>CO2EmissionKlasseNummer"`

	// KoeretoejMotor.xsd
	KoeretoejMotorCylinderAntal 		  		  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorCylinderAntal"`
	KoeretoejMotorKilometerstand 		  		  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorKilometerstand"`
	KoeretoejMotorKilometerstandDokumentation 		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorKilometerstandDokumentation"`
	KoeretoejMotorKilometerstandIkkeTilgaengelig 	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorKilometerstandIkkeTilgaengelig"`
	KoeretoejMotorKmPerLiter 		  		  	      			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorKmPerLiter"`
	KoeretoejMotorKMPerLiterPreCalc 		  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorKMPerLiterPreCalc"`
	KoeretoejMotorPlugInHybrid 		  		  	      			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorPlugInHybrid"`
	KoeretoejMotorKoerselStoej 		  		  	  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorKoerselStoej"`
	KoeretoejMotorMaerkning 		  		  	  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorMaerkning"`
	KoeretoejMotorSlagVolumen 		  		  	      			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorSlagVolumen"`
	KoeretoejMotorSlagVolumenIkkeTilgaengelig 		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorSlagVolumenIkkeTilgaengelig"`
	KoeretoejMotorStandStoej 		  		  	      			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorStandStoej"`
	KoeretoejMotorStandStoejOmdrejningstal 		  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorStandStoejOmdrejningstal"`
	KoeretoejMotorStoersteEffekt 		  		  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorStoersteEffekt"`
	KoeretoejMotorStoersteEffektIkkeTilgaengelig 	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorStoersteEffektIkkeTilgaengelig"`
	KoeretoejMotorInnovativTeknik 		  		  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorInnovativTeknik"`
	KoeretoejMotorInnovativTeknikAntal 		  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorInnovativTeknikAntal"`
	KoeretoejMotorElektriskForbrug 		  		  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorElektriskForbrug"`
	KoeretoejMotorFuelmode 		  		  	          			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorFuelmode"`
	KoeretoejMotorGasforbrug 		  		  	      			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorGasforbrug"`
	KoeretoejMotorElektriskRaekkevidde 		  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorElektriskRaekkevidde"`
	KoeretoejMotorBatterikapacitet 		  		  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorBatterikapacitet"`
	KoeretoejMotorBraendstofforbrugMaalt 		  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorBraendstofforbrugMaalt"`
	KoeretoejMotorElektriskForbrugMaalt 		  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorElektriskForbrugMaalt"`
	KoeretoejMotorMaaleNormTypeNavn 		  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorMaaleNormTypeNavn"`
	KoeretoejMotorMaaleNormTypeNummer 		  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorMaaleNormTypeNummer"`
	KoeretoejMotorCO2UdslipBeregnet 		  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorCO2UdslipBeregnet"`
	KoeretoejMotorBraendselscelle 		  		  	  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorBraendselscelle"`
	KoeretoejMotorDrivmiddelPrimaer 		  		  			  string `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejMotorDrivmiddelPrimaer"`

	// KoeretoejDrivmiddelSamlingStrukturType.xsd
	KoeretoejDrivmiddelSamling 					      			  KoeretoejDrivmiddelSamlingStruct `xml:"KoeretoejOplysningGrundStruktur>KoeretoejMotorStruktur>KoeretoejDrivmiddelSamlingStruktur>KoeretoejDrivmiddelSamling"`

	// DispensationTypeSamlingStrukturType.xsd
	DispensationTypeSamling 					      			  DispensationTypeSamlingStruct `xml:"KoeretoejOplysningGrundStruktur>DispensationTypeSamlingStruktur"`

	// KoeretoejUdstyrSamlingStrukturType.xsd
	KoeretoejUdstyrSamling 			  				  			  KoeretoejUdstyrSamlingStruct `xml:"KoeretoejOplysningGrundStruktur>KoeretoejUdstyrSamlingStruktur>KoeretoejUdstyrSamling"`

	// KoeretoejBlokeringAArsagListeStrukturType.xsd
	KoeretoejBlokeringAarsagListe 					  			  KoeretoejBlokeringAarsagListeStruct `xml:"KoeretoejOplysningGrundStruktur>KoeretoejBlokeringAarsagListeStruktur>KoeretoejBlokeringAarsagListe"`

	// PrisOplysninger.xsd
	PrisOplysningerMindsteBeskatningspris             			  string `xml:"PrisOplysningerStruktur>PrisOplysningerMindsteBeskatningspris"`
	PrisOplysningerIndkoebsPris                  	  			  string `xml:"PrisOplysningerStruktur>PrisOplysningerIndkoebsPris"`
	PrisOplysningerStandardPris                  	  			  string `xml:"PrisOplysningerStruktur>PrisOplysningerStandardPris"`

	// KoeretoejGruppe.xsd
	KoeretoejGruppeNavn                               			  string `xml:"KoeretoejGruppeStruktur>KoeretoejGruppeNavn"`
	KoeretoejGruppeNummer                  	          			  string `xml:"KoeretoejGruppeStruktur>KoeretoejGruppeNummer"`

	// KoeretoejUndergruppe.xsd
	KoeretoejUndergruppeNavn                          			  string `xml:"KoeretoejUndergruppeStruktur>KoeretoejUndergruppeNavn"`
	KoeretoejUndergruppeNummer                  	  			  string `xml:"KoeretoejUndergruppeStruktur>KoeretoejUndergruppeNummer"`

	/* FUTURE WORKS: EjerBrugerSamlingStrukturType.xsd
	EjerBrugerSamling 					              			  EjerBrugerSamlingStruct `xml:"EjerBrugerSamling"`
	*/

	// Adresse.xsd
	AdresseFortloebendeNummer                   	  			  string `xml:"AdresseFortloebendeNummer"`
	AdresseAnvendelseKode                   	      			  string `xml:"AdresseAnvendelseKode"`
	AdresseVejNavn                   	              			  string `xml:"AdresseVejNavn"`
	AdresseVejKode                   	              			  string `xml:"AdresseVejKode"`
	AdresseFraHusNummer                   	          			  string `xml:"AdresseFraHusNummer"`
	AdresseFraHusBogstav                   	          			  string `xml:"AdresseFraHusBogstav"`
	AdresseTilHusNummer                   	          			  string `xml:"AdresseTilHusNummer"`
	AdresseTilHusBogstav                   	          			  string `xml:"AdresseTilHusBogstav"`
	AdresseLigeUlige                   	              			  string `xml:"AdresseLigeUlige"`
	AdresseLejlighedNummer                   	      			  string `xml:"AdresseLejlighedNummer"`
	AdresseHusNavn                   	              			  string `xml:"AdresseHusNavn"`
	AdresseEtage                   	                  			  string `xml:"AdresseEtage"`
	AdresseEtageTekst                   	          			  string `xml:"AdresseEtageTekst"`
	AdresseSideDoerTekst                   	          			  string `xml:"AdresseSideDoerTekst"`
	AdresseCONavn                   	              			  string `xml:"AdresseCONavn"`
	AdressePostNummer                  	              			  string `xml:"AdressePostNummer"`
	AdressePostDistrikt                   	          			  string `xml:"AdressePostDistrikt"`
	AdresseLandsBy                   	              			  string `xml:"AdresseLandsBy"`
	AdresseByNavn                   	              			  string `xml:"AdresseByNavn"`
	AdresseLandsDel                   	              			  string `xml:"AdresseLandsDel"`
	AdressePostBox                   	              			  string `xml:"AdressePostBox"`
	AdresseGyldigFra                   	              			  string `xml:"AdresseGyldigFra"`
	AdresseGyldigTil                   	              			  string `xml:"AdresseGyldigTil"`

	// SynResultat.xsd
	SynResultatNummer                  			      			  string `xml:"SynResultatStruktur>SynResultatNummer"`
	SynResultatSynsDato                  			  			  string `xml:"SynResultatStruktur>SynResultatSynsDato"`
	SynResultatSynsResultat                  		  			  string `xml:"SynResultatStruktur>SynResultatSynsResultat"`
	SynResultatSynStatus                  			  			  string `xml:"SynResultatStruktur>SynResultatSynStatus"`
	SynResultatSynStatusDato                  		  			  string `xml:"SynResultatStruktur>SynResultatSynStatusDato"`
	SynResultatSynsType                  			  			  string `xml:"SynResultatStruktur>SynResultatSynsType"`
	SynResultatOmsynMoedeDato                  		  			  string `xml:"SynResultatStruktur>SynResultatOmsynMoedeDato"`
	SynResultatKoeretoejMotorKilometerstand           			  string `xml:"SynResultatStruktur>KoeretoejMotorKilometerstand"`

	/* FUTURE WORKS: KoeretoejRegistreringGrundlagSamlingStrukturType.xsd
	*/

	// KoeretoejRegistrering.xsd
	KoeretoejRegistreringGyldigFra                    			  string `xml:"KoeretoejRegistreringGyldigFra"`
	KoeretoejRegistreringGyldigTil                    			  string `xml:"KoeretoejRegistreringGyldigTil"`
	KoeretoejRegistreringNummer                       			  string `xml:"KoeretoejRegistreringNummer"`
	KoeretoejRegistreringStatus                  	  			  string `xml:"KoeretoejRegistreringStatus"`
	KoeretoejRegistreringStatusDato                   			  string `xml:"KoeretoejRegistreringStatusDato"`
	KoeretoejRegistreringStatusAarsag                 			  string `xml:"KoeretoejRegistreringStatusAarsag"`
	KoeretoejRegistreringKontrolTal                   			  string `xml:"KoeretoejRegistreringKontrolTal"`
	KoeretoejRegistreringGrundlagIdent                			  string `xml:"KoeretoejRegistreringGrundlagIdent"`
	KoeretoejRegistreringSenesteHaendelse             			  string `xml:"KoeretoejRegistreringSenesteHaendelse"`
	KoeretoejRegistreringTilknyttetLeasingForhold     			  string `xml:"KoeretoejRegistreringTilknyttetLeasingForhold"`

	// TilladelseSamlingStrukturType.xsd
	Tilladelse 			  				  			  			  TilladelseStruct `xml:"TilladelseSamling>Tilladelse"`
}

func main() {

	var total int64

	// Start the timer
	startTime := time.Now()

	workerCount := runtime.NumCPU()
  	const batchSize = 1000

	// Open the ZIP file
	zipFile, err := zip.OpenReader("ESStatistikListeModtag-20250420-165437.zip")
	if err != nil {
		fmt.Println("Error opening ZIP file:", err)
		return
	}
	defer zipFile.Close()

	// Find the XML file inside the ZIP file
	var xmlFile *zip.File
	for _, file := range zipFile.File {
		if file.Name == "ESStatistikListeModtag.xml" {
			xmlFile = file
			break
		}
	}

	if xmlFile == nil {

		fmt.Println("XML file not found in the ZIP file")
		return
	}

	// Open the XML file inside the ZIP file
	xmlFileReader, err := xmlFile.Open()
	if err != nil {


		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFileReader.Close()

	// Create an XML decoder
	decoder := xml.NewDecoder(xmlFileReader)

	// statistikCount := 0

	// Open a connection to MySQL database
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/bilgaden_new_new")
	db.SetMaxOpenConns(workerCount * 2)
	db.SetMaxIdleConns(workerCount * 2)
	db.SetConnMaxLifetime(time.Hour)
	if err != nil {


		fmt.Println("Error connecting to MySQL:", err)
		return
	}
	defer db.Close()

	koeretoejStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoej (
			KoeretoejIdent
		) VALUES (
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {

		fmt.Println("Error preparing SQL koeretoejStatement:", err)
		return
	}
	defer koeretoejStatement.Close()

	koeretoejAnvendelseStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejanvendelse (
			KoeretoejIdent,
			KoeretoejAnvendelseNummer,
			KoeretoejAnvendelseNavn,
			KoeretoejAnvendelseBeskrivelse,
			KoeretoejAnvendelseGyldigFra,
			KoeretoejAnvendelseGyldigTil,
			KoeretoejAnvendelseStatus
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {

		fmt.Println("Error preparing SQL koeretoejAnvendelseStatement:", err)
		return
	}
	defer koeretoejAnvendelseStatement.Close()

	koeretoejAnvendelseSupplerendeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejanvendelsesupplerende (
			KoeretoejIdent,
			KoeretoejAnvendelseSupplerendeNummer,
			KoeretoejAnvendelseSupplerendeNavn,
			KoeretoejAnvendelseSupplerendeBeskrivelse,
			KoeretoejAnvendelseSupplerendeGyldigFra,
			KoeretoejAnvendelseSupplerendeGyldigTil,
			KoeretoejAnvendelseSupplerendeStatus
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejAnvendelseSupplerendeStatement:", err)
		return
	}
	defer koeretoejAnvendelseSupplerendeStatement.Close()

	koeretoejArtStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejart (
			KoeretoejIdent,
			KoeretoejArtNummer,
			KoeretoejArtNavn,
			KoeretoejArtKraeverForsikring,
			KoeretoejArtBeskrivelse,
			KoeretoejArtGyldigFra,
			KoeretoejArtGyldigTil,
			KoeretoejArtStatus
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejArtStatement:", err)
		return
	}
	defer koeretoejArtStatement.Close()

	koeretoejBlokeringAarsagStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejblokeringaarsag (
			KoeretoejIdent,
			KoeretoejBlokeringAarsagTypeNummer,
			KoeretoejBlokeringAarsagTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejBlokeringAarsagStatement:", err)
		return
	}
	defer koeretoejBlokeringAarsagStatement.Close()

	koeretoejBraendstofStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejbraendstof (
			KoeretoejIdent,
			KoeretoejMotorKmPerLiter,
			KoeretoejMotorKMPerLiterPreCalc,
			KoeretoejMotorBraendstofforbrugMaalt,
			KoeretoejMotorGasforbrug,
			KoeretoejMotorCO2UdslipBeregnet,
			KoeretoejMiljoeOplysningCO2Udslip
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejBraendstofStatement:", err)
		return
	}
	defer koeretoejBraendstofStatement.Close()

	koeretoejCO2EmissionKlasseStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejco2emissionklasse (
			KoeretoejIdent,
			CO2EmissionKlasseNummer,
			CO2EmissionKlasseNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejCO2EmissionKlasseStatement:", err)
		return
	}
	defer koeretoejCO2EmissionKlasseStatement.Close()

	koeretoejDispensationTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejdispensationtype (
			KoeretoejIdent,
			DispensationTypeNummer,
			DispensationTypeNavn,
			KoeretoejDispensationTypeKommentar
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejDispensationTypeStatement:", err)
		return
	}
	defer koeretoejDispensationTypeStatement.Close()

	koeretoejDrivkraftTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejdrivkrafttype (
			KoeretoejIdent,
			DrivkraftTypeNummer,
			DrivkraftTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejDrivkraftTypeStatement:", err)
		return
	}
	defer koeretoejDrivkraftTypeStatement.Close()

	koeretoejDrivmiddelStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejdrivmiddel (
			KoeretoejIdent,
			KoeretoejMotorBraendselscelle,
			KoeretoejMotorPlugInHybrid,
			KoeretoejMotorDrivmiddelPrimaer
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejDrivmiddelStatement:", err)
		return
	}
	defer koeretoejDrivmiddelStatement.Close()

	koeretoejElforbrugStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejelforbrug (
			KoeretoejIdent,
			KoeretoejMotorElektriskForbrug,
			KoeretoejMotorElektriskForbrugMaalt,
			KoeretoejMotorElektriskRaekkevidde,
			KoeretoejMotorBatterikapacitet
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejElforbrugStatement:", err)
		return
	}
	defer koeretoejElforbrugStatement.Close()

	koeretoejFarveTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejfarvetype (
			KoeretoejIdent,
			FarveTypeNummer,
			FarveTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejFarveTypeStatement:", err)
		return
	}
	defer koeretoejFarveTypeStatement.Close()

	koeretoejFastKombinationStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejfastkombination (
			KoeretoejIdent,
			KoeretoejFastKombinationIdent,
			RegistreringNummerNummer,
			RegistreringNummerIdent
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejFastKombinationStatement:", err)
		return
	}
	defer koeretoejFastKombinationStatement.Close()

	koeretoejGruppeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejgruppe (
			KoeretoejIdent,
			KoeretoejGruppeNummer,
			KoeretoejGruppeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejGruppeStatement:", err)
		return
	}
	defer koeretoejGruppeStatement.Close()

	koeretoejKarrosseriTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejkarrosseritype (
			KoeretoejIdent,
			KarrosseriTypeNummer,
			KarrosseriTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejKarrosseriTypeStatement:", err)
		return
	}
	defer koeretoejKarrosseriTypeStatement.Close()

	koeretoejLeasingStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejleasing (
			KoeretoejIdent,
			LeasingMaaneder,
			LeasingNummer,
			LeasingGyldigFra,
			LeasingGyldigTil,
			LeasingReelOphoerDato,
			LeasingKode,
			LeasingStatus,
			LeasingBemaerkning,
			LeasingAendringType,
			LeasingSidstAendret
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejLeasingStatement:", err)
		return
	}
	defer koeretoejLeasingStatement.Close()

	koeretoejMaaleNormStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejmaalenorm (
			KoeretoejIdent,
			KoeretoejMotorMaaleNormTypeNummer,
			KoeretoejMotorMaaleNormTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejMaaleNormStatement:", err)
		return
	}
	defer koeretoejMaaleNormStatement.Close()

	koeretoejMaerkeTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejmaerketype (
			KoeretoejIdent,
			KoeretoejMaerkeTypeNummer,
			KoeretoejMaerkeTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejMaerkeTypeStatement:", err)
		return
	}
	defer koeretoejMaerkeTypeStatement.Close()

	koeretoejMiljoeOplysningStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejmiljoeoplysning (
			KoeretoejIdent,
			KoeretoejMiljoeOplysningCO2EmissionKlasseOmklassificeringDato,
			KoeretoejMiljoeOplysningCO2EmissionKlasseOprindelig,
			KoeretoejMiljoeOplysningCO2Udslip,
			KoeretoejMiljoeOplysningEftermonteretPartikelfilter,
			KoeretoejMiljoeOplysningEmissionCO,
			KoeretoejMiljoeOplysningEmissionHCPlusNOX,
			KoeretoejMiljoeOplysningEmissionNOX,
			KoeretoejMiljoeOplysningNyttelastvaerdi,
			KoeretoejMiljoeOplysningPartikelFilter,
			KoeretoejMiljoeOplysningPartikler,
			KoeretoejMiljoeOplysningRoegtaethed,
			KoeretoejMiljoeOplysningRoegtaethedOmdrejningstal,
			KoeretoejMiljoeOplysningSpecifikCO2Emission,
			KoeretoejMiljoeOplysningTungtNulEmissionKoeretoej
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejMiljoeOplysningStatement:", err)
		return
	}
	defer koeretoejMiljoeOplysningStatement.Close()

	koeretoejModelTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejmodeltype (
			KoeretoejIdent,
			KoeretoejModelTypeNummer,
			KoeretoejModelTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejModelTypeStatement:", err)
		return
	}
	defer koeretoejModelTypeStatement.Close()

	koeretoejMotorStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejmotor (
			KoeretoejIdent,
			KoeretoejMotorCylinderAntal,
			KoeretoejMotorSlagVolumen,
			KoeretoejMotorSlagVolumenIkkeTilgaengelig,
			KoeretoejMotorStoersteEffekt,
			KoeretoejMotorStoersteEffektIkkeTilgaengelig,
			KoeretoejMotorKilometerstand,
			KoeretoejMotorKilometerstandDokumentation,
			KoeretoejMotorKilometerstandIkkeTilgaengelig,
			KoeretoejMotorKmPerLiter,
			KoeretoejMotorKMPerLiterPreCalc,
			KoeretoejMotorPlugInHybrid,
			KoeretoejMotorMaerkning,
			KoeretoejMotorStandStoej,
			KoeretoejMotorKoerselStoej,
			KoeretoejMotorStandStoejOmdrejningstal,
			KoeretoejMotorInnovativTeknik,
			KoeretoejMotorInnovativTeknikAntal,
			KoeretoejMotorElektriskForbrug,
			KoeretoejMotorFuelmode,
			KoeretoejMotorGasforbrug,
			KoeretoejMotorElektriskRaekkevidde,
			KoeretoejMotorBatterikapacitet,
			KoeretoejMotorBraendstofforbrugMaalt,
			KoeretoejMotorElektriskForbrugMaalt,
			KoeretoejMotorMaaleNormTypeNavn,
			KoeretoejMotorMaaleNormTypeNummer,
			KoeretoejMotorCO2UdslipBeregnet,
			KoeretoejMotorBraendselscelle,
			KoeretoejMotorDrivmiddelPrimaer
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejMotorStatement:", err)
		return
	}
	defer koeretoejMotorStatement.Close()

	koeretoejNormTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejnormtype (
			KoeretoejIdent,
			NormTypeNummer,
			NormTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejNormTypeStatement:", err)
		return
	}
	defer koeretoejNormTypeStatement.Close()

	koeretoejOplysningStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejoplysning (
			KoeretoejIdent,
			KoeretoejOplysningOprettetUdFra,
			KoeretoejOplysningStatus,
			KoeretoejOplysningStatusDato,
			KoeretoejOplysningFoersteRegistreringDato,
			KoeretoejOplysningStelNummer,
			KoeretoejOplysningStelNummerAnbringelse,
			KoeretoejOplysningModelAar,
			KoeretoejOplysningTotalVaegt,
			KoeretoejOplysningEgenVaegt,
			KoeretoejOplysningKoereklarVaegtMinimum,
			KoeretoejOplysningKoereklarVaegtMaksimum,
			KoeretoejOplysningTekniskTotalVaegt,
			KoeretoejOplysningVogntogVaegt,
			KoeretoejOplysningAkselAntal,
			KoeretoejOplysningStoersteAkselTryk,
			KoeretoejOplysningSkatteAkselAntal,
			KoeretoejOplysningSkatteAkselTryk,
			KoeretoejOplysningPassagerAntal,
			KoeretoejOplysningSiddepladserMinimum,
			KoeretoejOplysningSiddepladserMaksimum,
			KoeretoejOplysningStaapladserMinimum,
			KoeretoejOplysningStaapladserMaksimum,
			KoeretoejOplysningTilkoblingMulighed,
			KoeretoejOplysningTilkoblingsvaegtUdenBremser,
			KoeretoejOplysningTilkoblingsvaegtMedBremser,
			KoeretoejOplysningPaahaengVognTotalVaegt,
			KoeretoejOplysningSkammelBelastning,
			KoeretoejOplysningSaettevognTilladtAkselTryk,
			KoeretoejOplysningMaksimumHastighed,
			KoeretoejOplysningFaelgDaek,
			KoeretoejOplysningTilkobletSidevognStelnr,
			KoeretoejOplysningNCAPTest,
			KoeretoejOplysningVVaerdiLuft,
			KoeretoejOplysningVVaerdiMekanisk,
			KoeretoejOplysningOevrigtUdstyr,
			KoeretoejOplysningKoeretoejstand,
			KoeretoejOplysning30PctVarevogn,
			KoeretoejOplysningBlokvognAkselType,
			KoeretoejOplysningBlokvognHovedboltTryk,
			KoeretoejOplysningBlokvognSkammelTryk,
			KoeretoejOplysningBlokvognSamletAkselTryk,
			KoeretoejOplysningBlokvognMaxVogntog,
			KoeretoejOplysningBlokvognBreddeFra,
			KoeretoejOplysningBlokvognKoblingshoejdeFra,
			KoeretoejOplysningBlokvognKoblingslaengdeFra,
			KoeretoejOplysningBlokvognSammenkoblingType,
			KoeretoejOplysningBlokvognTilladeligHastighed,
			KoeretoejOplysningBlokvognBreddeTil,
			KoeretoejOplysningBlokvognKoblingshoejdeTil,
			KoeretoejOplysningBlokvognKoblingslaengdeTil,
			KoeretoejOplysningTraekkendeAksler,
			KoeretoejOplysningEgnetTilTaxi,
			KoeretoejOplysningAkselAfstand,
			KoeretoejOplysningSporviddenForrest,
			KoeretoejOplysningSporviddenBagest,
			KoeretoejOplysningTypeAnmeldelseNummer,
			KoeretoejOplysningTypeGodkendelseNummer,
			KoeretoejOplysningEUVariant,
			KoeretoejOplysningEUVersion,
			KoeretoejOplysningKommentar,
			KoeretoejOplysningTypegodkendtKategori,
			KoeretoejOplysningAntalGear,
			KoeretoejOplysningAntalDoere,
			KoeretoejOplysningFabrikant,
			KoeretoejOplysningFrikoert,
			KoeretoejOplysningFredetForPladeInddragelse,
			KoeretoejOplysningVejvenligLuftaffjedring,
			KoeretoejOplysningDanskGodkendelseNummer,
			KoeretoejOplysningAargang,
			KoeretoejOplysningIbrugtagningDato,
			KoeretoejOplysningTrafikskade,
			KoeretoejOplysningVeteranKoeretoejOriginal,
			KoeretoejOplysningEffektivitetforholdRelevant,
			KoeretoejOplysningEffektivitetforholdM3,
			KoeretoejOplysningEffektivitetforholdTon,
			KoeretoejOplysningVolumenorientering,
			KoeretoejOplysningSovekabine
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejOplysningStatement:", err)
		return
	}
	defer koeretoejOplysningStatement.Close()

	koeretoejPrisOplysningerStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejprisoplysninger (
			KoeretoejIdent,
			PrisOplysningerStandardPris,
			PrisOplysningerIndkoebsPris,
			PrisOplysningerMindsteBeskatningspris
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejPrisOplysningerStatement:", err)
		return
	}
	defer koeretoejPrisOplysningerStatement.Close()

	koeretoejRegistreringStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejregistrering (
			KoeretoejIdent,
			KoeretoejRegistreringNummer,
			KoeretoejRegistreringStatus,
			KoeretoejRegistreringStatusDato,
			KoeretoejRegistreringStatusAarsag,
			KoeretoejRegistreringKontrolTal,
			KoeretoejRegistreringGyldigFra,
			KoeretoejRegistreringGyldigTil,
			KoeretoejRegistreringGrundlagIdent,
			KoeretoejRegistreringSenesteHaendelse,
			KoeretoejRegistreringTilknyttetLeasingForhold
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejRegistreringStatement:", err)
		return
	}
	defer koeretoejRegistreringStatement.Close()

	koeretoejRegistreringGrundlagStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejregistreringgrundlag (
			KoeretoejIdent,
			KoeretoejRegistreringGrundlagIdent,
			KoeretoejRegistreringGrundlagStatus,
			KoeretoejRegistreringGrundlagStatusDato,
			KoeretoejRegistreringGrundlagGyldigFra,
			KoeretoejRegistreringGrundlagGyldigTil,
			KoeretoejRegistreringGrundlagKode,
			KoeretoejRegistreringGrundlagTilknyttetFasteProeveskilte,
			KoeretoejRegistreringGrundlagPeriodiskSyn,
			KoeretoejRegistreringGrundlagPeriodiskSynGyldigTil
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejRegistreringGrundlagStatement:", err)
		return
	}
	defer koeretoejRegistreringGrundlagStatement.Close()

	koeretoejRegistreringGrundlagAnvendelseStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejregistreringgrundlaganvendelse (
			KoeretoejIdent,
			KoeretoejAnvendelseNummer,
			KoeretoejAnvendelseNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejRegistreringGrundlagAnvendelseStatement:", err)
		return
	}
	defer koeretoejRegistreringGrundlagAnvendelseStatement.Close()

	koeretoejRegistreringGrundlagArtStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejregistreringgrundlagart (
			KoeretoejIdent,
			KoeretoejArtNummer,
			KoeretoejArtNavn,
			KoeretoejArtKraeverForsikring,
			KoeretoejArtBeskrivelse,
			KoeretoejArtGyldigFra,
			KoeretoejArtGyldigTil,
			KoeretoejArtStatus
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejRegistreringGrundlagArtStatement:", err)
		return
	}
	defer koeretoejRegistreringGrundlagArtStatement.Close()

	koeretoejRegistreringGrundlagGenerelIdentifikatorStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejregistreringgrundlaggenerelidentifikator (
			KoeretoejIdent,
			RegistreringNummerNummer,
			KoeretoejOplysningStelNummer
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejRegistreringGrundlagGenerelIdentifikatorStatement:", err)
		return
	}
	defer koeretoejRegistreringGrundlagGenerelIdentifikatorStatement.Close()

	koeretoejRegistreringNummerStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejregistreringnummer (
			KoeretoejIdent,
			RegistreringNummerIdent,
			RegistreringNummerType,
			RegistreringNummerStatus,
			RegistreringNummerStatusDato,
			RegistreringNummerKvadratiskIndhold1,
			RegistreringNummerKvadratiskIndhold2,
			RegistreringNummerAflangIndhold,
			RegistreringNummerUdloebDato,
			RegistreringNummerFigurantPlade,
			RegistreringNummerGraensepladeDkDato,
			RegistreringNummerNummer
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejRegistreringNummerStatement:", err)
		return
	}
	defer koeretoejRegistreringNummerStatement.Close()

	koeretoejRegistreringNummerRettighedStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejregistreringnummerrettighed (
			KoeretoejIdent,
			RegistreringNummerRettighedType,
			RegistreringNummerRettighedGyldigFra,
			RegistreringNummerRettighedGyldigTil,
			RegistreringNummerRettighedNummer,
			RegistreringNummerRettighedSidstAdviseretDato,
			RegistreringNummerRettighedKoerselFormaal,
			RegistreringNummerRettighedAntalFerieDageTilbage
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejRegistreringNummerRettighedStatement:", err)
		return
	}
	defer koeretoejRegistreringNummerRettighedStatement.Close()

	koeretoejSupplerendeKarrosseriTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejsupplerendekarrosseritype (
			KoeretoejIdent,
			SupplerendeKarrosseriTypeNummer,
			SupplerendeKarrosseriTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejSupplerendeKarrosseriTypeStatement:", err)
		return
	}
	defer koeretoejSupplerendeKarrosseriTypeStatement.Close()

	koeretoejSynResultatStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejsynresultat (
			KoeretoejIdent,
			KoeretoejMotorKilometerstand,
			SynResultatNummer,
			SynResultatSynStatus,
			SynResultatSynStatusDato,
			SynResultatSynsDato,
			SynResultatSynsResultat,
			SynResultatSynsType,
			SynResultatOmsynMoedeDato
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {

		fmt.Println("Error preparing SQL koeretoejSynResultatStatement:", err)
		return
	}
	defer koeretoejSynResultatStatement.Close()

	koeretoejTilladelseStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtilladelse (
			KoeretoejIdent,
			TilladelseNummer,
			TilladelseGyldigFra,
			TilladelseGyldigTil,
			TilladelseKommentar,
			TilladelseKunGodkendtForRegistreretEjer,
			TilladelseKombinationKoeretoejIdent
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejTilladelseStatement:", err)
		return
	}
	defer koeretoejTilladelseStatement.Close()

	koeretoejTilladelseTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtilladelsetype (
			KoeretoejIdent,
			TilladelseTypeNummer,
			TilladelseTypeNavn,
			TilladelseTypeErPeriodeBegraenset,
			TilladelseTypePeriodeLaengde
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejTilladelseTypeStatement:", err)
		return
	}
	defer koeretoejTilladelseTypeStatement.Close()

	koeretoejTilladelseTypeFastTilkoblingStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtilladelsetypefasttilkobling (
			KoeretoejIdent,
			KoeretoejFastTilkoblingIdent,
			KoeretoejOplysningStelNummer
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejTilladelseTypeFastTilkoblingStatement:", err)
		return
	}
	defer koeretoejTilladelseTypeFastTilkoblingStatement.Close()

	koeretoejTilladelseTypeKunGodkendtForJuridiskEnhedStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtilladelsetypekungodkendtforjuridiskenhed (
			KoeretoejIdent,
			PersonCPRNummer,
			VirksomhedSENummer,
			VirksomhedCVRNummer,
			ProduktionEnhedNummer,
			AlternativKontaktID
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {

		fmt.Println("Error preparing SQL koeretoejTilladelseTypeKunGodkendtForJuridiskEnhedStatement:", err)
		return
	}
	defer koeretoejTilladelseTypeKunGodkendtForJuridiskEnhedStatement.Close()

	koeretoejTilladelseTypeVariabelKombinationStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtilladelsetypevariabelkombination (
			KoeretoejIdent,
			KoeretoejVariabelKombinationIdent,
			RegistreringNummerNummer,
			KoeretoejOplysningStelNummer
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejTilladelseTypeVariabelKombinationStatement:", err)
		return
	}
	defer koeretoejTilladelseTypeVariabelKombinationStatement.Close()

	koeretoejTypeAttestStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattest (
			KoeretoejIdent,
			TypeAttestGyldigFra,
			TypeAttestGyldigTil,
			TypeAttestTypeGodkendelseNummer,
			TypeAttestTypeAnmeldelseNummer
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejTypeAttestStatement:", err)
		return
	}
	defer koeretoejTypeAttestStatement.Close()

	koeretoejTypeAttestKoeretoejArtStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattestkoeretoejart (
			KoeretoejIdent,
			KoeretoejArtNummer,
			KoeretoejArtNavn,
			KoeretoejArtKraeverForsikring,
			KoeretoejArtBeskrivelse,
			KoeretoejArtGyldigFra,
			KoeretoejArtGyldigTil,
			KoeretoejArtStatus
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejTypeAttestKoeretoejArtStatement:", err)
		return
	}
	defer koeretoejTypeAttestKoeretoejArtStatement.Close()

	koeretoejTypeAttestKoeretoejMaerkeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattestkoeretoejmaerke (
			KoeretoejIdent,
			KoeretoejMaerkeTypeNummer,
			KoeretoejMaerkeTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejTypeAttestKoeretoejMaerkeStatement:", err)
		return
	}
	defer koeretoejTypeAttestKoeretoejMaerkeStatement.Close()

	koeretoejTypeAttestKoeretoejModelStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattestkoeretoejmodel (
			KoeretoejIdent,
			KoeretoejModelTypeNummer,
			KoeretoejModelTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejTypeAttestKoeretoejModelStatement:", err)
		return
	}
	defer koeretoejTypeAttestKoeretoejModelStatement.Close()

	koeretoejTypeAttestKoeretoejTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattestkoeretoejtype (
			KoeretoejIdent,
			KoeretoejTypeTypeNummer,
			KoeretoejTypeTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejTypeAttestKoeretoejTypeStatement:", err)
		return
	}
	defer koeretoejTypeAttestKoeretoejTypeStatement.Close()

	koeretoejTypeAttestKoeretoejVariantStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattestkoeretoejvariant (
			KoeretoejIdent,
			KoeretoejVariantTypeNummer,
			KoeretoejVariantTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejTypeAttestKoeretoejVariantStatement:", err)
		return
	}
	defer koeretoejTypeAttestKoeretoejVariantStatement.Close()

	koeretoejTypeAttestTilladelseStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattesttilladelse (
			KoeretoejIdent,
			TilladelseTypeNummer,
			TilladelseTypeNavn,
			TilladelseTypeErPeriodeBegraenset,
			TilladelseTypePeriodeLaengde
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejTypeAttestTilladelseStatement:", err)
		return
	}
	defer koeretoejTypeAttestTilladelseStatement.Close()

	koeretoejTypeAttestVariantStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattestvariant (
			KoeretoejIdent,
			TypeAttestVariantNummer,
			TypeAttestVariantSiddepladserMinimum,
			TypeAttestVariantSiddepladserMaksimum,
			TypeAttestVariantEgenVaegt,
			TypeAttestVariantKoereklarVaegtMaksimum,
			TypeAttestVariantTekniskTotalVaegt,
			TypeAttestVariantTotalVaegt,
			TypeAttestVariantStoersteAkselTryk,
			TypeAttestVariantTilkoblingsvaegtMedBremser,
			TypeAttestVariantTilkoblingsvaegtUdenBremser,
			TypeAttestVariantStatus,
			TypeAttestVariantStatusDatoTid,
			TypeAttestVariantStaapladserMinimum,
			TypeAttestVariantStaapladserMaksimum,
			TypeAttestVariantPassagerAntal,
			TypeAttestVariantAkselAntal,
			TypeAttestVariantFaelgDaek,
			TypeAttestVariantMaksimumHastighed,
			TypeAttestVariantStelNummerAnbringelse,
			TypeAttestVariantVVaerdiLuft,
			TypeAttestVariantVVaerdiMekanisk,
			TypeAttestVariantTraekkendeAksler,
			TypeAttestVariantAntalGear,
			TypeAttestVariantAntalDoere,
			TypeAttestVariantCO2Udslip,
			TypeAttestVariantRoegtaethed,
			TypeAttestVariantRoegtaethedOmdrejningstal,
			TypeAttestVariantPartikelFilter,
			TypeAttestVariantCylinderAntal,
			TypeAttestVariantMaerkning,
			TypeAttestVariantStandStoej,
			TypeAttestVariantStandStoejOmdrejningstal,
			TypeAttestVariantKoerselStoej,
			TypeAttestVariantKoereklarVaegtMinimum,
			TypeAttestVariantEgnetTilTaxi,
			TypeAttestVariantPartikler,
			TypeAttestVariantKmPerLiter,
			TypeAttestVariantStoersteEffekt,
			TypeAttestVariantInnovativTeknik,
			TypeAttestVariantInnovativTeknikAntal,
			TypeAttestVariantNCAPTest,
			TypeAttestVariantSkammelBelastning,
			TypeAttestVariantSkatteAkselAntal,
			TypeAttestVariantSkatteAkselTryk,
			TypeAttestVariantSaettevognTilladtAkselTryk,
			TypeAttestVariantVogntogVaegt,
			TypeAttestVariantAkselAfstand,
			TypeAttestVariantSporviddenForrest,
			TypeAttestVariantSporviddenBagest,
			TypeAttestVariantSlagVolumen,
			TypeAttestVariantElektriskForbrug,
			KoeretoejVariantTypeNummer,
			KoeretoejTypeTypeNummer
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejTypeAttestVariantStatement:", err)
		return
	}
	defer koeretoejTypeAttestVariantStatement.Close()

	koeretoejTypeAttestVariantDrivkraftStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattestvariantdrivkraft (
			KoeretoejIdent,
			DrivkraftTypeNummer,
			DrivkraftTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejTypeAttestVariantDrivkraftStatement:", err)
		return
	}
	defer koeretoejTypeAttestVariantDrivkraftStatement.Close()

	koeretoejTypeAttestVariantKarrosseriStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattestvariantkarrosseri (
			KoeretoejIdent,
			KarrosseriTypeNummer,
			KarrosseriTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejTypeAttestVariantKarrosseriStatement:", err)
		return
	}
	defer koeretoejTypeAttestVariantKarrosseriStatement.Close()

	koeretoejTypeAttestVariantNormStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypeattestvariantnorm (
			KoeretoejIdent,
			NormTypeNummer,
			NormTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejTypeAttestVariantNormStatement:", err)
		return
	}
	defer koeretoejTypeAttestVariantNormStatement.Close()

	koeretoejTypeTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejtypetype (
			KoeretoejIdent,
			KoeretoejTypeTypeNummer,
			KoeretoejTypeTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejTypeTypeStatement:", err)
		return
	}
	defer koeretoejTypeTypeStatement.Close()

	koeretoejUdstyrStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejudstyr (
			KoeretoejIdent,
			KoeretoejUdstyrAntal,
			KoeretoejUdstyrTypeNummer,
			KoeretoejUdstyrTypeNavn,
			KoeretoejUdstyrTypeVisesVedSyn,
			KoeretoejUdstyrTypeVisesVedForespoergsel,
			KoeretoejUdstyrTypeVisesVedStandardOprettelse,
			KoeretoejUdstyrTypeStandardAntal
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejUdstyrStatement:", err)
		return
	}
	defer koeretoejUdstyrStatement.Close()

	koeretoejUndergruppeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejundergruppe (
			KoeretoejIdent,
			KoeretoejUndergruppeNummer,
			KoeretoejUndergruppeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {


		fmt.Println("Error preparing SQL koeretoejUndergruppeStatement:", err)
		return
	}
	defer koeretoejUndergruppeStatement.Close()

	koeretoejVariantTypeStatement, err := db.Prepare(`
		INSERT IGNORE INTO koeretoejvarianttype (
			KoeretoejIdent,
			KoeretoejVariantTypeNummer,
			KoeretoejVariantTypeNavn
		) VALUES (
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL'),
			NULLIF(?, 'NULL')
		)
	`)
	if err != nil {



		fmt.Println("Error preparing SQL koeretoejVariantTypeStatement:", err)
		return
	}
	defer koeretoejVariantTypeStatement.Close()

  // 3) Build a cancelable context so any worker error can stop everyone
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  // 4) Channel + waitgroup
  statistikCh := make(chan Statistik, workerCount*2)
  var wg sync.WaitGroup
  wg.Add(workerCount)

  // 5) Spawn workers
  for i := 0; i < workerCount; i++ {
	  go func(workerID int) {
		  defer wg.Done()

		  // Each worker keeps its own tx + batch counter
		  tx, err := db.BeginTx(ctx, nil)
		  if err != nil {
			  log.Printf("[worker %d] begin tx: %v", workerID, err)
			  cancel()
			  return
		  }
		  inBatch := 0

		  // Ensure we clean up
		  defer func() {
			  if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
				  log.Printf("[worker %d] rollback: %v", workerID, err)
			  }
		  }()

		  for {
			  select {
			  case <-ctx.Done():
				  return
			  case statistik, ok := <-statistikCh:
				  if !ok {
					  // no more work
					  if inBatch > 0 {
						  if err := tx.Commit(); err != nil {
							  log.Printf("[worker %d] final commit: %v", workerID, err)
							  cancel()
						  }
					  }
					  return
				  }

				  // --- your per‑row work goes here ---
				  statistik.KoeretoejIdent = trimAndSetEmptyToNull(statistik.KoeretoejIdent)
				  if (statistik.KoeretoejIdent != "NULL") {
							  // bind the global stmt into this tx
							if _, err := tx.StmtContext(ctx, koeretoejStatement).
								ExecContext(ctx, statistik.KoeretoejIdent); err != nil {
								fmt.Println("Error executing SQL koeretoejStatement:", err)
							}

							  statistik.KoeretoejArtNavn = trimAndSetEmptyToNull(statistik.KoeretoejArtNavn) //YC
							  statistik.KoeretoejArtNummer = trimAndSetEmptyToNull(statistik.KoeretoejArtNummer) //YC
							  statistik.KoeretoejArtKraeverForsikring = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejArtKraeverForsikring)) //NC
							  statistik.KoeretoejArtBeskrivelse = trimAndSetEmptyToNull(statistik.KoeretoejArtBeskrivelse) //NC
							  statistik.KoeretoejArtGyldigFra = trimAndSetEmptyToNull(statistik.KoeretoejArtGyldigFra) //NC
							  statistik.KoeretoejArtGyldigTil = trimAndSetEmptyToNull(statistik.KoeretoejArtGyldigTil) //NC
							  statistik.KoeretoejArtStatus = trimAndSetEmptyToNull(statistik.KoeretoejArtStatus) //NC

							  if (statistik.KoeretoejArtNavn != "NULL" ||
								  statistik.KoeretoejArtNummer != "NULL" ||
									 statistik.KoeretoejArtKraeverForsikring != "NULL" ||
									 statistik.KoeretoejArtBeskrivelse != "NULL" ||
									 statistik.KoeretoejArtGyldigFra != "NULL" ||
									 statistik.KoeretoejArtGyldigTil != "NULL" ||
									 statistik.KoeretoejArtStatus != "NULL") {

								// bind the global stmt into this tx
								_, err := tx.StmtContext(ctx, koeretoejArtStatement).
									ExecContext(ctx,
									statistik.KoeretoejIdent,

									  statistik.KoeretoejArtNummer,
									  statistik.KoeretoejArtNavn,
									  statistik.KoeretoejArtKraeverForsikring,
									  statistik.KoeretoejArtBeskrivelse,
									  statistik.KoeretoejArtGyldigFra,
									  statistik.KoeretoejArtGyldigTil,
									  statistik.KoeretoejArtStatus,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejArtStatement:", err)
								  }
							  }

							  statistik.KoeretoejAnvendelseNavn = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseNavn) //YC
							  statistik.KoeretoejAnvendelseNummer = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseNummer) //YC
							  statistik.KoeretoejAnvendelseBeskrivelse = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseBeskrivelse) //NC
							  statistik.KoeretoejAnvendelseGyldigFra = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseGyldigFra) //NC
							  statistik.KoeretoejAnvendelseGyldigTil = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseGyldigTil) //NC
							  statistik.KoeretoejAnvendelseStatus = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseStatus) //NC

							  if (statistik.KoeretoejAnvendelseNavn != "NULL" ||
								  statistik.KoeretoejAnvendelseNummer != "NULL" ||
								  statistik.KoeretoejAnvendelseBeskrivelse != "NULL" ||
								  statistik.KoeretoejAnvendelseGyldigFra != "NULL" ||
								  statistik.KoeretoejAnvendelseGyldigTil != "NULL" ||
								  statistik.KoeretoejAnvendelseStatus != "NULL") {

								_, err := tx.StmtContext(ctx, koeretoejAnvendelseStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejAnvendelseNummer,
									  statistik.KoeretoejAnvendelseNavn,
									  statistik.KoeretoejAnvendelseBeskrivelse,
									  statistik.KoeretoejAnvendelseGyldigFra,
									  statistik.KoeretoejAnvendelseGyldigTil,
									  statistik.KoeretoejAnvendelseStatus,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejAnvendelseStatement:", err)
								  }
							  }

							  // FOR-LOOP FOR KoeretoejAnvendelseSamling
							  for i := range statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse {
								  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseNavn = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseNavn) //YC
								  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseNummer = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseNummer) //YC
								  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseBeskrivelse = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseBeskrivelse) //NC
								  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseGyldigFra = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseGyldigFra) //NC
								  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseGyldigTil = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseGyldigTil) //NC
								  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseStatus = trimAndSetEmptyToNull(statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseStatus) //NC

								  if (statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseNavn != "NULL" ||
									  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseNummer != "NULL" ||
									  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseBeskrivelse != "NULL" ||
									  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseGyldigFra != "NULL" ||
									  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseGyldigTil != "NULL" ||
									  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseStatus != "NULL") {

									_, err := tx.StmtContext(ctx, koeretoejAnvendelseStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseNummer,
										  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseNavn,
										  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseBeskrivelse,
										  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseGyldigFra,
										  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseGyldigTil,
										  statistik.KoeretoejAnvendelseSamling.KoeretoejAnvendelse[i].KoeretoejAnvendelseStatus,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejAnvendelseSupplerendeStatement in for-loop:", err)
									  }
								  }
							  }

							  statistik.LeasingMaaneder = trimAndSetEmptyToNull(statistik.LeasingMaaneder) //NC
							  statistik.LeasingNummer = trimAndSetEmptyToNull(statistik.LeasingNummer) //NC
							  statistik.LeasingGyldigFra = trimAndSetEmptyToNull(statistik.LeasingGyldigFra) //YC
							  statistik.LeasingGyldigTil = trimAndSetEmptyToNull(statistik.LeasingGyldigTil) //YC
							  statistik.LeasingReelOphoerDato = trimAndSetEmptyToNull(statistik.LeasingReelOphoerDato) //NC
							  statistik.LeasingKode = trimAndSetEmptyToNull(statistik.LeasingKode) //NC
							  statistik.LeasingStatus = trimAndSetEmptyToNull(statistik.LeasingStatus) //NC
							  statistik.LeasingBemaerkning = trimAndSetEmptyToNull(statistik.LeasingBemaerkning) //NC
							  statistik.LeasingAendringType = trimAndSetEmptyToNull(statistik.LeasingAendringType) //NC
							  statistik.LeasingSidstAendret = trimAndSetEmptyToNull(statistik.LeasingSidstAendret) //NC

							  if (statistik.LeasingMaaneder != "NULL" ||
								  statistik.LeasingNummer != "NULL" ||
								  statistik.LeasingGyldigFra != "NULL" ||
								  statistik.LeasingGyldigTil != "NULL" ||
								  statistik.LeasingReelOphoerDato != "NULL" ||
								  statistik.LeasingKode != "NULL" ||
								  statistik.LeasingStatus != "NULL" ||
								  statistik.LeasingBemaerkning != "NULL" ||
								  statistik.LeasingAendringType != "NULL" ||
								  statistik.LeasingSidstAendret != "NULL") {

								_, err := tx.StmtContext(ctx, koeretoejLeasingStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.LeasingMaaneder,
									  statistik.LeasingNummer,
									  statistik.LeasingGyldigFra,
									  statistik.LeasingGyldigTil,
									  statistik.LeasingReelOphoerDato,
									  statistik.LeasingKode,
									  statistik.LeasingStatus,
									  statistik.LeasingBemaerkning,
									  statistik.LeasingAendringType,
									  statistik.LeasingSidstAendret,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejLeasingStatement:", err)
								  }
							  }

							  statistik.RegistreringNummerIdent = trimAndSetEmptyToNull(statistik.RegistreringNummerIdent) //NC
							  statistik.RegistreringNummerAflangIndhold = trimAndSetEmptyToNull(statistik.RegistreringNummerAflangIndhold) //NC
							  statistik.RegistreringNummerGraensepladeDkDato = trimAndSetEmptyToNull(statistik.RegistreringNummerGraensepladeDkDato) //NC
							  statistik.RegistreringNummerKvadratiskIndhold1 = trimAndSetEmptyToNull(statistik.RegistreringNummerKvadratiskIndhold1) //NC
							  statistik.RegistreringNummerKvadratiskIndhold2 = trimAndSetEmptyToNull(statistik.RegistreringNummerKvadratiskIndhold2) //NC
							  statistik.RegistreringNummerNummer = trimAndSetEmptyToNull(statistik.RegistreringNummerNummer) //YC
							  statistik.RegistreringNummerStatus = trimAndSetEmptyToNull(statistik.RegistreringNummerStatus) //NC
							  statistik.RegistreringNummerStatusDato = trimAndSetEmptyToNull(statistik.RegistreringNummerStatusDato) //NC
							  statistik.RegistreringNummerType = trimAndSetEmptyToNull(statistik.RegistreringNummerType) //NC
							  statistik.RegistreringNummerUdloebDato = trimAndSetEmptyToNull(statistik.RegistreringNummerUdloebDato) //YC
							  statistik.RegistreringNummerFigurantPlade = matchTrueFalse(trimAndSetEmptyToNull(statistik.RegistreringNummerFigurantPlade)) //NC

							  if (statistik.RegistreringNummerIdent != "NULL" ||
								  statistik.RegistreringNummerAflangIndhold != "NULL" ||
								  statistik.RegistreringNummerGraensepladeDkDato != "NULL" ||
								  statistik.RegistreringNummerKvadratiskIndhold1 != "NULL" ||
								  statistik.RegistreringNummerKvadratiskIndhold2 != "NULL" ||
								  statistik.RegistreringNummerNummer != "NULL" ||
								  statistik.RegistreringNummerStatus != "NULL" ||
								  statistik.RegistreringNummerStatusDato != "NULL" ||
								  statistik.RegistreringNummerType != "NULL" ||
								  statistik.RegistreringNummerUdloebDato != "NULL" ||
								  statistik.RegistreringNummerFigurantPlade != "NULL") {

								_, err := tx.StmtContext(ctx, koeretoejRegistreringNummerStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.RegistreringNummerIdent,
									  statistik.RegistreringNummerType,
									  statistik.RegistreringNummerStatus,
									  statistik.RegistreringNummerStatusDato,
									  statistik.RegistreringNummerKvadratiskIndhold1,
									  statistik.RegistreringNummerKvadratiskIndhold2,
									  statistik.RegistreringNummerAflangIndhold,
									  statistik.RegistreringNummerUdloebDato,
									  statistik.RegistreringNummerFigurantPlade,
									  statistik.RegistreringNummerGraensepladeDkDato,
									  statistik.RegistreringNummerNummer,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejRegistreringNummerStatement:", err)
								  }
							  }

							  statistik.RegistreringNummerRettighedGyldigFra = trimAndSetEmptyToNull(statistik.RegistreringNummerRettighedGyldigFra) //YC
							  statistik.RegistreringNummerRettighedGyldigTil = trimAndSetEmptyToNull(statistik.RegistreringNummerRettighedGyldigTil) //YC
							  statistik.RegistreringNummerRettighedNummer = trimAndSetEmptyToNull(statistik.RegistreringNummerRettighedNummer) //NC
							  statistik.RegistreringNummerRettighedSidstAdviseretDato = trimAndSetEmptyToNull(statistik.RegistreringNummerRettighedSidstAdviseretDato) //NC
							  statistik.RegistreringNummerRettighedType = trimAndSetEmptyToNull(statistik.RegistreringNummerRettighedType) //NC
							  statistik.RegistreringNummerRettighedKoerselFormaal = trimAndSetEmptyToNull(statistik.RegistreringNummerRettighedKoerselFormaal) //NC
							  statistik.RegistreringNummerRettighedAntalFerieDageTilbage = trimAndSetEmptyToNull(statistik.RegistreringNummerRettighedAntalFerieDageTilbage) //NC

							  if (statistik.RegistreringNummerRettighedGyldigFra != "NULL" ||
								  statistik.RegistreringNummerRettighedGyldigTil != "NULL" ||
								  statistik.RegistreringNummerRettighedNummer != "NULL" ||
								  statistik.RegistreringNummerRettighedSidstAdviseretDato != "NULL" ||
								  statistik.RegistreringNummerRettighedType != "NULL" ||
								  statistik.RegistreringNummerRettighedKoerselFormaal != "NULL" ||
								  statistik.RegistreringNummerRettighedAntalFerieDageTilbage != "NULL") {

								_, err := tx.StmtContext(ctx, koeretoejRegistreringNummerRettighedStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.RegistreringNummerRettighedType,
									  statistik.RegistreringNummerRettighedGyldigFra,
									  statistik.RegistreringNummerRettighedGyldigTil,
									  statistik.RegistreringNummerRettighedNummer,
									  statistik.RegistreringNummerRettighedSidstAdviseretDato,
									  statistik.RegistreringNummerRettighedKoerselFormaal,
									  statistik.RegistreringNummerRettighedAntalFerieDageTilbage,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejRegistreringNummerRettighedStatement:", err)
								  }
							  }

							  statistik.KoeretoejOplysningOprettetUdFra = trimAndSetEmptyToNull(statistik.KoeretoejOplysningOprettetUdFra) //YC
							  statistik.KoeretoejOplysningStatus = trimAndSetEmptyToNull(statistik.KoeretoejOplysningStatus) //YC
							  statistik.KoeretoejOplysningStatusDato = trimAndSetEmptyToNull(statistik.KoeretoejOplysningStatusDato) //YC
							  statistik.KoeretoejOplysningFoersteRegistreringDato = trimAndSetEmptyToNull(statistik.KoeretoejOplysningFoersteRegistreringDato) //YC
							  statistik.KoeretoejOplysningStelNummer = trimAndSetEmptyToNull(statistik.KoeretoejOplysningStelNummer) //YC
							  statistik.KoeretoejOplysningStelNummerAnbringelse = trimAndSetEmptyToNull(statistik.KoeretoejOplysningStelNummerAnbringelse) //YC
							  statistik.KoeretoejOplysningModelAar = trimAndSetEmptyToNull(statistik.KoeretoejOplysningModelAar) //YC
							  statistik.KoeretoejOplysningTotalVaegt = trimAndSetEmptyToNull(statistik.KoeretoejOplysningTotalVaegt) //YC
							  statistik.KoeretoejOplysningEgenVaegt = trimAndSetEmptyToNull(statistik.KoeretoejOplysningEgenVaegt) //YC
							  statistik.KoeretoejOplysningKoereklarVaegtMinimum = trimAndSetEmptyToNull(statistik.KoeretoejOplysningKoereklarVaegtMinimum) //YC
							  statistik.KoeretoejOplysningKoereklarVaegtMaksimum = trimAndSetEmptyToNull(statistik.KoeretoejOplysningKoereklarVaegtMaksimum) //YC
							  statistik.KoeretoejOplysningTekniskTotalVaegt = trimAndSetEmptyToNull(statistik.KoeretoejOplysningTekniskTotalVaegt) //YC
							  statistik.KoeretoejOplysningVogntogVaegt = trimAndSetEmptyToNull(statistik.KoeretoejOplysningVogntogVaegt) //YC
							  statistik.KoeretoejOplysningAkselAntal = trimAndSetEmptyToNull(statistik.KoeretoejOplysningAkselAntal) //YC
							  statistik.KoeretoejOplysningStoersteAkselTryk = trimAndSetEmptyToNull(statistik.KoeretoejOplysningStoersteAkselTryk) //YC
							  statistik.KoeretoejOplysningSkatteAkselAntal = trimAndSetEmptyToNull(statistik.KoeretoejOplysningSkatteAkselAntal) //YC
							  statistik.KoeretoejOplysningSkatteAkselTryk = trimAndSetEmptyToNull(statistik.KoeretoejOplysningSkatteAkselTryk) //YC
							  statistik.KoeretoejOplysningPassagerAntal = trimAndSetEmptyToNull(statistik.KoeretoejOplysningPassagerAntal) //YC
							  statistik.KoeretoejOplysningSiddepladserMinimum = trimAndSetEmptyToNull(statistik.KoeretoejOplysningSiddepladserMinimum) //YC
							  statistik.KoeretoejOplysningSiddepladserMaksimum = trimAndSetEmptyToNull(statistik.KoeretoejOplysningSiddepladserMaksimum) //YC
							  statistik.KoeretoejOplysningStaapladserMinimum = trimAndSetEmptyToNull(statistik.KoeretoejOplysningStaapladserMinimum) //YC
							  statistik.KoeretoejOplysningStaapladserMaksimum = trimAndSetEmptyToNull(statistik.KoeretoejOplysningStaapladserMaksimum) //YC
							  statistik.KoeretoejOplysningTilkoblingMulighed = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningTilkoblingMulighed)) //YC
							  statistik.KoeretoejOplysningTilkoblingsvaegtUdenBremser = trimAndSetEmptyToNull(statistik.KoeretoejOplysningTilkoblingsvaegtUdenBremser) //YC
							  statistik.KoeretoejOplysningTilkoblingsvaegtMedBremser = trimAndSetEmptyToNull(statistik.KoeretoejOplysningTilkoblingsvaegtMedBremser) //YC
							  statistik.KoeretoejOplysningPaahaengVognTotalVaegt = trimAndSetEmptyToNull(statistik.KoeretoejOplysningPaahaengVognTotalVaegt) //YC
							  statistik.KoeretoejOplysningSkammelBelastning = trimAndSetEmptyToNull(statistik.KoeretoejOplysningSkammelBelastning) //YC
							  statistik.KoeretoejOplysningSaettevognTilladtAkselTryk = trimAndSetEmptyToNull(statistik.KoeretoejOplysningSaettevognTilladtAkselTryk) //YC
							  statistik.KoeretoejOplysningMaksimumHastighed = trimAndSetEmptyToNull(statistik.KoeretoejOplysningMaksimumHastighed) //YC
							  statistik.KoeretoejOplysningFaelgDaek = trimAndSetEmptyToNull(statistik.KoeretoejOplysningFaelgDaek) //YC
							  statistik.KoeretoejOplysningTilkobletSidevognStelnr = trimAndSetEmptyToNull(statistik.KoeretoejOplysningTilkobletSidevognStelnr) //YC
							  statistik.KoeretoejOplysningNCAPTest = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningNCAPTest)) //YC
							  statistik.KoeretoejOplysningVVaerdiLuft = trimAndSetEmptyToNull(statistik.KoeretoejOplysningVVaerdiLuft) //YC
							  statistik.KoeretoejOplysningVVaerdiMekanisk = trimAndSetEmptyToNull(statistik.KoeretoejOplysningVVaerdiMekanisk) //YC
							  statistik.KoeretoejOplysningOevrigtUdstyr = trimAndSetEmptyToNull(statistik.KoeretoejOplysningOevrigtUdstyr) //YC
							  statistik.KoeretoejOplysningKoeretoejstand = trimAndSetEmptyToNull(statistik.KoeretoejOplysningKoeretoejstand) //YC
							  statistik.KoeretoejOplysning30PctVarevogn = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysning30PctVarevogn)) //YC
							  statistik.KoeretoejOplysningBlokvognAkselType = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognAkselType) //NC
							  statistik.KoeretoejOplysningBlokvognHovedboltTryk = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognHovedboltTryk) //NC
							  statistik.KoeretoejOplysningBlokvognSkammelTryk = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognSkammelTryk) //NC
							  statistik.KoeretoejOplysningBlokvognSamletAkselTryk = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognSamletAkselTryk) //NC
							  statistik.KoeretoejOplysningBlokvognMaxVogntog = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognMaxVogntog) //NC
							  statistik.KoeretoejOplysningBlokvognBreddeFra = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognBreddeFra) //NC
							  statistik.KoeretoejOplysningBlokvognKoblingshoejdeFra = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognKoblingshoejdeFra) //NC
							  statistik.KoeretoejOplysningBlokvognKoblingslaengdeFra = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognKoblingslaengdeFra) //NC
							  statistik.KoeretoejOplysningBlokvognSammenkoblingType = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognSammenkoblingType) //NC
							  statistik.KoeretoejOplysningBlokvognTilladeligHastighed = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognTilladeligHastighed) //NC
							  statistik.KoeretoejOplysningBlokvognBreddeTil = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognBreddeTil) //NC
							  statistik.KoeretoejOplysningBlokvognKoblingshoejdeTil = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognKoblingshoejdeTil) //NC
							  statistik.KoeretoejOplysningBlokvognKoblingslaengdeTil = trimAndSetEmptyToNull(statistik.KoeretoejOplysningBlokvognKoblingslaengdeTil) //NC
							  statistik.KoeretoejOplysningTraekkendeAksler = trimAndSetEmptyToNull(statistik.KoeretoejOplysningTraekkendeAksler) //YC
							  statistik.KoeretoejOplysningEgnetTilTaxi = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningEgnetTilTaxi)) //YC
							  statistik.KoeretoejOplysningAkselAfstand = trimAndSetEmptyToNull(statistik.KoeretoejOplysningAkselAfstand) //YC
							  statistik.KoeretoejOplysningSporviddenForrest = trimAndSetEmptyToNull(statistik.KoeretoejOplysningSporviddenForrest) //YC
							  statistik.KoeretoejOplysningSporviddenBagest = trimAndSetEmptyToNull(statistik.KoeretoejOplysningSporviddenBagest) //YC
							  statistik.KoeretoejOplysningTypeAnmeldelseNummer = trimAndSetEmptyToNull(statistik.KoeretoejOplysningTypeAnmeldelseNummer) //YC
							  statistik.KoeretoejOplysningTypeGodkendelseNummer = trimAndSetEmptyToNull(statistik.KoeretoejOplysningTypeGodkendelseNummer) //YC
							  statistik.KoeretoejOplysningEUVariant = trimAndSetEmptyToNull(statistik.KoeretoejOplysningEUVariant) //YC
							  statistik.KoeretoejOplysningEUVersion = trimAndSetEmptyToNull(statistik.KoeretoejOplysningEUVersion) //YC
							  statistik.KoeretoejOplysningKommentar = trimAndSetEmptyToNull(statistik.KoeretoejOplysningKommentar) //YC
							  statistik.KoeretoejOplysningTypegodkendtKategori = trimAndSetEmptyToNull(statistik.KoeretoejOplysningTypegodkendtKategori) //YC
							  statistik.KoeretoejOplysningAntalGear = trimAndSetEmptyToNull(statistik.KoeretoejOplysningAntalGear) //YC
							  statistik.KoeretoejOplysningAntalDoere = trimAndSetEmptyToNull(statistik.KoeretoejOplysningAntalDoere) //YC
							  statistik.KoeretoejOplysningFabrikant = trimAndSetEmptyToNull(statistik.KoeretoejOplysningFabrikant) //YC
							  statistik.KoeretoejOplysningFrikoert = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningFrikoert)) //NC
							  statistik.KoeretoejOplysningFredetForPladeInddragelse = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningFredetForPladeInddragelse)) //NC
							  statistik.KoeretoejOplysningVejvenligLuftaffjedring = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningVejvenligLuftaffjedring)) //NC
							  statistik.KoeretoejOplysningDanskGodkendelseNummer = trimAndSetEmptyToNull(statistik.KoeretoejOplysningDanskGodkendelseNummer) //NC
							  statistik.KoeretoejOplysningAargang = trimAndSetEmptyToNull(statistik.KoeretoejOplysningAargang) //NC
							  statistik.KoeretoejOplysningIbrugtagningDato = trimAndSetEmptyToNull(statistik.KoeretoejOplysningIbrugtagningDato) //YC
							  statistik.KoeretoejOplysningTrafikskade = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningTrafikskade)) //YC
							  statistik.KoeretoejOplysningVeteranKoeretoejOriginal = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningVeteranKoeretoejOriginal)) //YC
							  statistik.KoeretoejOplysningEffektivitetforholdRelevant = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningEffektivitetforholdRelevant)) //NC
							  statistik.KoeretoejOplysningEffektivitetforholdM3 = trimAndSetEmptyToNull(statistik.KoeretoejOplysningEffektivitetforholdM3) //NC
							  statistik.KoeretoejOplysningEffektivitetforholdTon = trimAndSetEmptyToNull(statistik.KoeretoejOplysningEffektivitetforholdTon) //NC
							  statistik.KoeretoejOplysningVolumenorientering = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningVolumenorientering)) //NC
							  statistik.KoeretoejOplysningSovekabine = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejOplysningSovekabine)) //NC

							  if (statistik.KoeretoejOplysningOprettetUdFra != "NULL" ||
								  statistik.KoeretoejOplysningStatus != "NULL" ||
								  statistik.KoeretoejOplysningStatusDato != "NULL" ||
								  statistik.KoeretoejOplysningFoersteRegistreringDato != "NULL" ||
								  statistik.KoeretoejOplysningStelNummer != "NULL" ||
								  statistik.KoeretoejOplysningStelNummerAnbringelse != "NULL" ||
								  statistik.KoeretoejOplysningModelAar != "NULL" ||
								  statistik.KoeretoejOplysningTotalVaegt != "NULL" ||
								  statistik.KoeretoejOplysningEgenVaegt != "NULL" ||
								  statistik.KoeretoejOplysningKoereklarVaegtMinimum != "NULL" ||
								  statistik.KoeretoejOplysningKoereklarVaegtMaksimum != "NULL" ||
								  statistik.KoeretoejOplysningTekniskTotalVaegt != "NULL" ||
								  statistik.KoeretoejOplysningVogntogVaegt != "NULL" ||
								  statistik.KoeretoejOplysningAkselAntal != "NULL" ||
								  statistik.KoeretoejOplysningStoersteAkselTryk != "NULL" ||
								  statistik.KoeretoejOplysningSkatteAkselAntal != "NULL" ||
								  statistik.KoeretoejOplysningSkatteAkselTryk != "NULL" ||
								  statistik.KoeretoejOplysningPassagerAntal != "NULL" ||
								  statistik.KoeretoejOplysningSiddepladserMinimum != "NULL" ||
								  statistik.KoeretoejOplysningSiddepladserMaksimum != "NULL" ||
								  statistik.KoeretoejOplysningStaapladserMinimum != "NULL" ||
								  statistik.KoeretoejOplysningStaapladserMaksimum != "NULL" ||
								  statistik.KoeretoejOplysningTilkoblingMulighed != "NULL" ||
								  statistik.KoeretoejOplysningTilkoblingsvaegtUdenBremser != "NULL" ||
								  statistik.KoeretoejOplysningTilkoblingsvaegtMedBremser != "NULL" ||
								  statistik.KoeretoejOplysningPaahaengVognTotalVaegt != "NULL" ||
								  statistik.KoeretoejOplysningSkammelBelastning != "NULL" ||
								  statistik.KoeretoejOplysningSaettevognTilladtAkselTryk != "NULL" ||
								  statistik.KoeretoejOplysningMaksimumHastighed != "NULL" ||
								  statistik.KoeretoejOplysningFaelgDaek != "NULL" ||
								  statistik.KoeretoejOplysningTilkobletSidevognStelnr != "NULL" ||
								  statistik.KoeretoejOplysningNCAPTest != "NULL" ||
								  statistik.KoeretoejOplysningVVaerdiLuft != "NULL" ||
								  statistik.KoeretoejOplysningVVaerdiMekanisk != "NULL" ||
								  statistik.KoeretoejOplysningOevrigtUdstyr != "NULL" ||
								  statistik.KoeretoejOplysningKoeretoejstand != "NULL" ||
								  statistik.KoeretoejOplysning30PctVarevogn != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognAkselType != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognHovedboltTryk != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognSkammelTryk != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognSamletAkselTryk != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognMaxVogntog != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognBreddeFra != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognKoblingshoejdeFra != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognKoblingslaengdeFra != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognSammenkoblingType != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognTilladeligHastighed != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognBreddeTil != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognKoblingshoejdeTil != "NULL" ||
								  statistik.KoeretoejOplysningBlokvognKoblingslaengdeTil != "NULL" ||
								  statistik.KoeretoejOplysningTraekkendeAksler != "NULL" ||
								  statistik.KoeretoejOplysningEgnetTilTaxi != "NULL" ||
								  statistik.KoeretoejOplysningAkselAfstand != "NULL" ||
								  statistik.KoeretoejOplysningSporviddenForrest != "NULL" ||
								  statistik.KoeretoejOplysningSporviddenBagest != "NULL" ||
								  statistik.KoeretoejOplysningTypeAnmeldelseNummer != "NULL" ||
								  statistik.KoeretoejOplysningTypeGodkendelseNummer != "NULL" ||
								  statistik.KoeretoejOplysningEUVariant != "NULL" ||
								  statistik.KoeretoejOplysningEUVersion != "NULL" ||
								  statistik.KoeretoejOplysningKommentar != "NULL" ||
								  statistik.KoeretoejOplysningTypegodkendtKategori != "NULL" ||
								  statistik.KoeretoejOplysningAntalGear != "NULL" ||
								  statistik.KoeretoejOplysningAntalDoere != "NULL" ||
								  statistik.KoeretoejOplysningFabrikant != "NULL" ||
								  statistik.KoeretoejOplysningFrikoert != "NULL" ||
								  statistik.KoeretoejOplysningFredetForPladeInddragelse != "NULL" ||
								  statistik.KoeretoejOplysningVejvenligLuftaffjedring != "NULL" ||
								  statistik.KoeretoejOplysningDanskGodkendelseNummer != "NULL" ||
								  statistik.KoeretoejOplysningAargang != "NULL" ||
								  statistik.KoeretoejOplysningIbrugtagningDato != "NULL" ||
								  statistik.KoeretoejOplysningTrafikskade != "NULL" ||
								  statistik.KoeretoejOplysningVeteranKoeretoejOriginal != "NULL" ||
								  statistik.KoeretoejOplysningEffektivitetforholdRelevant != "NULL" ||
								  statistik.KoeretoejOplysningEffektivitetforholdM3 != "NULL" ||
								  statistik.KoeretoejOplysningEffektivitetforholdTon != "NULL" ||
								  statistik.KoeretoejOplysningVolumenorientering != "NULL" ||
								  statistik.KoeretoejOplysningSovekabine != "NULL") {

								_, err := tx.StmtContext(ctx, koeretoejOplysningStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejOplysningOprettetUdFra,
									  statistik.KoeretoejOplysningStatus,
									  statistik.KoeretoejOplysningStatusDato,
									  statistik.KoeretoejOplysningFoersteRegistreringDato,
									  statistik.KoeretoejOplysningStelNummer,
									  statistik.KoeretoejOplysningStelNummerAnbringelse,
									  statistik.KoeretoejOplysningModelAar,
									  statistik.KoeretoejOplysningTotalVaegt,
									  statistik.KoeretoejOplysningEgenVaegt,
									  statistik.KoeretoejOplysningKoereklarVaegtMinimum,
									  statistik.KoeretoejOplysningKoereklarVaegtMaksimum,
									  statistik.KoeretoejOplysningTekniskTotalVaegt,
									  statistik.KoeretoejOplysningVogntogVaegt,
									  statistik.KoeretoejOplysningAkselAntal,
									  statistik.KoeretoejOplysningStoersteAkselTryk,
									  statistik.KoeretoejOplysningSkatteAkselAntal,
									  statistik.KoeretoejOplysningSkatteAkselTryk,
									  statistik.KoeretoejOplysningPassagerAntal,
									  statistik.KoeretoejOplysningSiddepladserMinimum,
									  statistik.KoeretoejOplysningSiddepladserMaksimum,
									  statistik.KoeretoejOplysningStaapladserMinimum,
									  statistik.KoeretoejOplysningStaapladserMaksimum,
									  statistik.KoeretoejOplysningTilkoblingMulighed,
									  statistik.KoeretoejOplysningTilkoblingsvaegtUdenBremser,
									  statistik.KoeretoejOplysningTilkoblingsvaegtMedBremser,
									  statistik.KoeretoejOplysningPaahaengVognTotalVaegt,
									  statistik.KoeretoejOplysningSkammelBelastning,
									  statistik.KoeretoejOplysningSaettevognTilladtAkselTryk,
									  statistik.KoeretoejOplysningMaksimumHastighed,
									  statistik.KoeretoejOplysningFaelgDaek,
									  statistik.KoeretoejOplysningTilkobletSidevognStelnr,
									  statistik.KoeretoejOplysningNCAPTest,
									  statistik.KoeretoejOplysningVVaerdiLuft,
									  statistik.KoeretoejOplysningVVaerdiMekanisk,
									  statistik.KoeretoejOplysningOevrigtUdstyr,
									  statistik.KoeretoejOplysningKoeretoejstand,
									  statistik.KoeretoejOplysning30PctVarevogn,
									  statistik.KoeretoejOplysningBlokvognAkselType,
									  statistik.KoeretoejOplysningBlokvognHovedboltTryk,
									  statistik.KoeretoejOplysningBlokvognSkammelTryk,
									  statistik.KoeretoejOplysningBlokvognSamletAkselTryk,
									  statistik.KoeretoejOplysningBlokvognMaxVogntog,
									  statistik.KoeretoejOplysningBlokvognBreddeFra,
									  statistik.KoeretoejOplysningBlokvognKoblingshoejdeFra,
									  statistik.KoeretoejOplysningBlokvognKoblingslaengdeFra,
									  statistik.KoeretoejOplysningBlokvognSammenkoblingType,
									  statistik.KoeretoejOplysningBlokvognTilladeligHastighed,
									  statistik.KoeretoejOplysningBlokvognBreddeTil,
									  statistik.KoeretoejOplysningBlokvognKoblingshoejdeTil,
									  statistik.KoeretoejOplysningBlokvognKoblingslaengdeTil,
									  statistik.KoeretoejOplysningTraekkendeAksler,
									  statistik.KoeretoejOplysningEgnetTilTaxi,
									  statistik.KoeretoejOplysningAkselAfstand,
									  statistik.KoeretoejOplysningSporviddenForrest,
									  statistik.KoeretoejOplysningSporviddenBagest,
									  statistik.KoeretoejOplysningTypeAnmeldelseNummer,
									  statistik.KoeretoejOplysningTypeGodkendelseNummer,
									  statistik.KoeretoejOplysningEUVariant,
									  statistik.KoeretoejOplysningEUVersion,
									  statistik.KoeretoejOplysningKommentar,
									  statistik.KoeretoejOplysningTypegodkendtKategori,
									  statistik.KoeretoejOplysningAntalGear,
									  statistik.KoeretoejOplysningAntalDoere,
									  statistik.KoeretoejOplysningFabrikant,
									  statistik.KoeretoejOplysningFrikoert,
									  statistik.KoeretoejOplysningFredetForPladeInddragelse,
									  statistik.KoeretoejOplysningVejvenligLuftaffjedring,
									  statistik.KoeretoejOplysningDanskGodkendelseNummer,
									  statistik.KoeretoejOplysningAargang,
									  statistik.KoeretoejOplysningIbrugtagningDato,
									  statistik.KoeretoejOplysningTrafikskade,
									  statistik.KoeretoejOplysningVeteranKoeretoejOriginal,
									  statistik.KoeretoejOplysningEffektivitetforholdRelevant,
									  statistik.KoeretoejOplysningEffektivitetforholdM3,
									  statistik.KoeretoejOplysningEffektivitetforholdTon,
									  statistik.KoeretoejOplysningVolumenorientering,
									  statistik.KoeretoejOplysningSovekabine,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejOplysningStatement:", err)
								  }
							  }

							  statistik.KoeretoejFastKombinationKoeretoejIdent = trimAndSetEmptyToNull(statistik.KoeretoejFastKombinationKoeretoejIdent) //M - gør om igen
							  statistik.KoeretoejFastKombinationRegistreringNummerNummer = trimAndSetEmptyToNull(statistik.KoeretoejFastKombinationRegistreringNummerNummer) //M - gør om igen
							  statistik.KoeretoejFastKombinationRegistreringNummerIdent = trimAndSetEmptyToNull(statistik.KoeretoejFastKombinationRegistreringNummerIdent) //M - gør om igen

							  if (statistik.KoeretoejFastKombinationKoeretoejIdent != "NULL" ||
								  statistik.KoeretoejFastKombinationRegistreringNummerNummer != "NULL" ||
								  statistik.KoeretoejFastKombinationRegistreringNummerIdent != "NULL") {

								  if statistik.KoeretoejIdent != statistik.KoeretoejFastKombinationKoeretoejIdent {
									  // fmt.Printf(
									  // 	"\nstatistik.KoeretoejIdent:\n%s\ndoes not equal statistik.KoeretoejFastKombinationKoeretoejIdent:\n%s\n",
									  // 	statistik.KoeretoejIdent,
									  // 	statistik.KoeretoejFastKombinationKoeretoejIdent,
									  // )
								  }

								_, err = tx.StmtContext(ctx, koeretoejFastKombinationStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejFastKombinationKoeretoejIdent,
									  statistik.KoeretoejFastKombinationRegistreringNummerNummer,
									  statistik.KoeretoejFastKombinationRegistreringNummerIdent,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejFastKombinationStatement:", err)
								  }
							  }

							  statistik.KoeretoejMaerkeTypeNavn = trimAndSetEmptyToNull(statistik.KoeretoejMaerkeTypeNavn) //YC
							  statistik.KoeretoejMaerkeTypeNummer = trimAndSetEmptyToNull(statistik.KoeretoejMaerkeTypeNummer) //YC

							  if (statistik.KoeretoejMaerkeTypeNavn != "NULL" ||
								  statistik.KoeretoejMaerkeTypeNummer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejMaerkeTypeStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejMaerkeTypeNummer,
									  statistik.KoeretoejMaerkeTypeNavn,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejMaerkeTypeStatement:", err)
								  }
							  }

							  statistik.KoeretoejModelTypeNavn = trimAndSetEmptyToNull(statistik.KoeretoejModelTypeNavn) //YC
							  statistik.KoeretoejModelTypeNummer = trimAndSetEmptyToNull(statistik.KoeretoejModelTypeNummer) //YC

							  if (statistik.KoeretoejModelTypeNavn != "NULL" ||
								  statistik.KoeretoejModelTypeNummer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejModelTypeStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejModelTypeNummer,
									  statistik.KoeretoejModelTypeNavn,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejModelTypeStatement:", err)
								  }
							  }

							  statistik.KoeretoejTypeTypeNavn = trimAndSetEmptyToNull(statistik.KoeretoejTypeTypeNavn) //YC
							  statistik.KoeretoejTypeTypeNummer = trimAndSetEmptyToNull(statistik.KoeretoejTypeTypeNummer) //YC

							  if (statistik.KoeretoejTypeTypeNavn != "NULL" ||
								  statistik.KoeretoejTypeTypeNummer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejTypeTypeStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejTypeTypeNummer,
									  statistik.KoeretoejTypeTypeNavn,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejTypeTypeStatement:", err)
								  }
							  }

							  statistik.KoeretoejVariantTypeNavn = trimAndSetEmptyToNull(statistik.KoeretoejVariantTypeNavn) //YC
							  statistik.KoeretoejVariantTypeNummer = trimAndSetEmptyToNull(statistik.KoeretoejVariantTypeNummer) //YC

							  if (statistik.KoeretoejVariantTypeNavn != "NULL" ||
								  statistik.KoeretoejVariantTypeNummer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejVariantTypeStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejVariantTypeNummer,
									  statistik.KoeretoejVariantTypeNavn,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejVariantTypeStatement:", err)
								  }
							  }

							  statistik.FarveTypeNavn = trimAndSetEmptyToNull(statistik.FarveTypeNavn) //YC
							  statistik.FarveTypeNummer = trimAndSetEmptyToNull(statistik.FarveTypeNummer) //YC

							  if (statistik.FarveTypeNavn != "NULL" ||
								  statistik.FarveTypeNummer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejFarveTypeStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.FarveTypeNummer,
									  statistik.FarveTypeNavn,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejFarveTypeStatement:", err)
								  }
							  }

							  statistik.KarrosseriTypeNavn = trimAndSetEmptyToNull(statistik.KarrosseriTypeNavn) //YC
							  statistik.KarrosseriTypeNummer = trimAndSetEmptyToNull(statistik.KarrosseriTypeNummer) //YC

							  if (statistik.KarrosseriTypeNavn != "NULL" ||
								  statistik.KarrosseriTypeNummer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejKarrosseriTypeStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KarrosseriTypeNummer,
									  statistik.KarrosseriTypeNavn,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejKarrosseriTypeStatement:", err)
								  }
							  }

							  // FOR-LOOP FOR KoeretoejSupplerendeKarrosseriSamling
							  for i := range statistik.KoeretoejSupplerendeKarrosseriSamling.KoeretoejSupplerendeKarrosseriTypeStruktur {
								  statistik.KoeretoejSupplerendeKarrosseriSamling.KoeretoejSupplerendeKarrosseriTypeStruktur[i].SupplerendeKarrosseriTypeNavn = trimAndSetEmptyToNull(statistik.KoeretoejSupplerendeKarrosseriSamling.KoeretoejSupplerendeKarrosseriTypeStruktur[i].SupplerendeKarrosseriTypeNavn) //YC
								  statistik.KoeretoejSupplerendeKarrosseriSamling.KoeretoejSupplerendeKarrosseriTypeStruktur[i].SupplerendeKarrosseriTypeNummer = trimAndSetEmptyToNull(statistik.KoeretoejSupplerendeKarrosseriSamling.KoeretoejSupplerendeKarrosseriTypeStruktur[i].SupplerendeKarrosseriTypeNummer) //YC

								  if (statistik.KoeretoejSupplerendeKarrosseriSamling.KoeretoejSupplerendeKarrosseriTypeStruktur[i].SupplerendeKarrosseriTypeNavn != "NULL" ||
									  statistik.KoeretoejSupplerendeKarrosseriSamling.KoeretoejSupplerendeKarrosseriTypeStruktur[i].SupplerendeKarrosseriTypeNummer != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejSupplerendeKarrosseriTypeStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.KoeretoejSupplerendeKarrosseriSamling.KoeretoejSupplerendeKarrosseriTypeStruktur[i].SupplerendeKarrosseriTypeNummer,
										  statistik.KoeretoejSupplerendeKarrosseriSamling.KoeretoejSupplerendeKarrosseriTypeStruktur[i].SupplerendeKarrosseriTypeNavn,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejSupplerendeKarrosseriTypeStatement in for-loop:", err)
									  }
								  }
							  }

							  statistik.NormTypeNavn = trimAndSetEmptyToNull(statistik.NormTypeNavn) //YC
							  statistik.NormTypeNummer = trimAndSetEmptyToNull(statistik.NormTypeNummer) //YC

							  if (statistik.NormTypeNavn != "NULL" ||
								  statistik.NormTypeNummer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejNormTypeStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.NormTypeNummer,
									  statistik.NormTypeNavn,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejNormTypeStatement:", err)
								  }
							  }

							  statistik.KoeretoejMiljoeOplysningCO2Udslip = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningCO2Udslip) //NC
							  statistik.KoeretoejMiljoeOplysningEmissionCO = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningEmissionCO) //YC
							  statistik.KoeretoejMiljoeOplysningEmissionHCPlusNOX = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningEmissionHCPlusNOX) //YC
							  statistik.KoeretoejMiljoeOplysningEmissionNOX = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningEmissionNOX) //YC
							  statistik.KoeretoejMiljoeOplysningPartikler = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningPartikler) //YC
							  statistik.KoeretoejMiljoeOplysningPartikelFilter = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningPartikelFilter)) //YC
							  statistik.KoeretoejMiljoeOplysningRoegtaethed = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningRoegtaethed) //YC
							  statistik.KoeretoejMiljoeOplysningRoegtaethedOmdrejningstal = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningRoegtaethedOmdrejningstal) //YC
							  statistik.KoeretoejMiljoeOplysningEftermonteretPartikelfilter = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningEftermonteretPartikelfilter)) //NC
							  statistik.KoeretoejMiljoeOplysningSpecifikCO2Emission = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningSpecifikCO2Emission) //YC
							  statistik.KoeretoejMiljoeOplysningNyttelastvaerdi = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningNyttelastvaerdi) //YC
							  statistik.KoeretoejMiljoeOplysningTungtNulEmissionKoeretoej = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningTungtNulEmissionKoeretoej)) //YC
							  statistik.KoeretoejMiljoeOplysningCO2EmissionKlasseOmklassificeringDato = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningCO2EmissionKlasseOmklassificeringDato) //YC
							  statistik.KoeretoejMiljoeOplysningCO2EmissionKlasseOprindelig = trimAndSetEmptyToNull(statistik.KoeretoejMiljoeOplysningCO2EmissionKlasseOprindelig) //NC

							  if (statistik.KoeretoejMiljoeOplysningCO2Udslip != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningEmissionCO != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningEmissionHCPlusNOX != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningEmissionNOX != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningPartikler != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningPartikelFilter != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningRoegtaethed != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningRoegtaethedOmdrejningstal != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningEftermonteretPartikelfilter != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningSpecifikCO2Emission != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningNyttelastvaerdi != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningTungtNulEmissionKoeretoej != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningCO2EmissionKlasseOmklassificeringDato != "NULL" ||
								  statistik.KoeretoejMiljoeOplysningCO2EmissionKlasseOprindelig != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejMiljoeOplysningStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejMiljoeOplysningCO2EmissionKlasseOmklassificeringDato,
									  statistik.KoeretoejMiljoeOplysningCO2EmissionKlasseOprindelig,
									  statistik.KoeretoejMiljoeOplysningCO2Udslip,
									  statistik.KoeretoejMiljoeOplysningEftermonteretPartikelfilter,
									  statistik.KoeretoejMiljoeOplysningEmissionCO,
									  statistik.KoeretoejMiljoeOplysningEmissionHCPlusNOX,
									  statistik.KoeretoejMiljoeOplysningEmissionNOX,
									  statistik.KoeretoejMiljoeOplysningNyttelastvaerdi,
									  statistik.KoeretoejMiljoeOplysningPartikelFilter,
									  statistik.KoeretoejMiljoeOplysningPartikler,
									  statistik.KoeretoejMiljoeOplysningRoegtaethed,
									  statistik.KoeretoejMiljoeOplysningRoegtaethedOmdrejningstal,
									  statistik.KoeretoejMiljoeOplysningSpecifikCO2Emission,
									  statistik.KoeretoejMiljoeOplysningTungtNulEmissionKoeretoej,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejMiljoeOplysningStatement:", err)
								  }
							  }

							  statistik.CO2EmissionKlasseNavn = trimAndSetEmptyToNull(statistik.CO2EmissionKlasseNavn) //YC
							  statistik.CO2EmissionKlasseNummer = trimAndSetEmptyToNull(statistik.CO2EmissionKlasseNummer) //YC

							  if (statistik.CO2EmissionKlasseNavn != "NULL" ||
								  statistik.CO2EmissionKlasseNummer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejCO2EmissionKlasseStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.CO2EmissionKlasseNummer,
									  statistik.CO2EmissionKlasseNavn,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejCO2EmissionKlasseStatement:", err)
								  }
							  }

							  statistik.KoeretoejMotorCylinderAntal = trimAndSetEmptyToNull(statistik.KoeretoejMotorCylinderAntal) //YC
							  statistik.KoeretoejMotorKilometerstand = trimAndSetEmptyToNull(statistik.KoeretoejMotorKilometerstand) //YC
							  statistik.KoeretoejMotorKilometerstandDokumentation = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMotorKilometerstandDokumentation)) //YC
							  statistik.KoeretoejMotorKilometerstandIkkeTilgaengelig = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMotorKilometerstandIkkeTilgaengelig)) //YC
							  statistik.KoeretoejMotorKmPerLiter = trimAndSetEmptyToNull(statistik.KoeretoejMotorKmPerLiter) //NC
							  statistik.KoeretoejMotorKMPerLiterPreCalc = trimAndSetEmptyToNull(statistik.KoeretoejMotorKMPerLiterPreCalc) //NC
							  statistik.KoeretoejMotorPlugInHybrid = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMotorPlugInHybrid)) //NC
							  statistik.KoeretoejMotorKoerselStoej = trimAndSetEmptyToNull(statistik.KoeretoejMotorKoerselStoej) //YC
							  statistik.KoeretoejMotorMaerkning = trimAndSetEmptyToNull(statistik.KoeretoejMotorMaerkning) //YC
							  statistik.KoeretoejMotorSlagVolumen = trimAndSetEmptyToNull(statistik.KoeretoejMotorSlagVolumen) //YC
							  statistik.KoeretoejMotorSlagVolumenIkkeTilgaengelig = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMotorSlagVolumenIkkeTilgaengelig)) //YC
							  statistik.KoeretoejMotorStandStoej = trimAndSetEmptyToNull(statistik.KoeretoejMotorStandStoej) //YC
							  statistik.KoeretoejMotorStandStoejOmdrejningstal = trimAndSetEmptyToNull(statistik.KoeretoejMotorStandStoejOmdrejningstal) //YC
							  statistik.KoeretoejMotorStoersteEffekt = trimAndSetEmptyToNull(statistik.KoeretoejMotorStoersteEffekt) //YC
							  statistik.KoeretoejMotorStoersteEffektIkkeTilgaengelig = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMotorStoersteEffektIkkeTilgaengelig)) //YC
							  statistik.KoeretoejMotorInnovativTeknik = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMotorInnovativTeknik)) //YC
							  statistik.KoeretoejMotorInnovativTeknikAntal = trimAndSetEmptyToNull(statistik.KoeretoejMotorInnovativTeknikAntal) //YC
							  statistik.KoeretoejMotorElektriskForbrug = trimAndSetEmptyToNull(statistik.KoeretoejMotorElektriskForbrug) //NC
							  statistik.KoeretoejMotorFuelmode = trimAndSetEmptyToNull(statistik.KoeretoejMotorFuelmode) //NC
							  statistik.KoeretoejMotorGasforbrug = trimAndSetEmptyToNull(statistik.KoeretoejMotorGasforbrug) //NC
							  statistik.KoeretoejMotorElektriskRaekkevidde = trimAndSetEmptyToNull(statistik.KoeretoejMotorElektriskRaekkevidde) //NC
							  statistik.KoeretoejMotorBatterikapacitet = trimAndSetEmptyToNull(statistik.KoeretoejMotorBatterikapacitet) //NC
							  statistik.KoeretoejMotorBraendstofforbrugMaalt = trimAndSetEmptyToNull(statistik.KoeretoejMotorBraendstofforbrugMaalt) //NC
							  statistik.KoeretoejMotorElektriskForbrugMaalt = trimAndSetEmptyToNull(statistik.KoeretoejMotorElektriskForbrugMaalt) //NC
							  statistik.KoeretoejMotorMaaleNormTypeNavn = trimAndSetEmptyToNull(statistik.KoeretoejMotorMaaleNormTypeNavn) //NC
							  statistik.KoeretoejMotorMaaleNormTypeNummer = trimAndSetEmptyToNull(statistik.KoeretoejMotorMaaleNormTypeNummer) //NC
							  statistik.KoeretoejMotorCO2UdslipBeregnet = trimAndSetEmptyToNull(statistik.KoeretoejMotorCO2UdslipBeregnet) //NC
							  statistik.KoeretoejMotorBraendselscelle = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMotorBraendselscelle)) //NC
							  statistik.KoeretoejMotorDrivmiddelPrimaer = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejMotorDrivmiddelPrimaer)) //NC

							  if (statistik.KoeretoejMotorCylinderAntal != "NULL" ||
								  statistik.KoeretoejMotorKilometerstand != "NULL" ||
								  statistik.KoeretoejMotorKilometerstandDokumentation != "NULL" ||
								  statistik.KoeretoejMotorKilometerstandIkkeTilgaengelig != "NULL" ||
								  statistik.KoeretoejMotorKmPerLiter != "NULL" ||
								  statistik.KoeretoejMotorKMPerLiterPreCalc != "NULL" ||
								  statistik.KoeretoejMotorPlugInHybrid != "NULL" ||
								  statistik.KoeretoejMotorKoerselStoej != "NULL" ||
								  statistik.KoeretoejMotorMaerkning != "NULL" ||
								  statistik.KoeretoejMotorSlagVolumen != "NULL" ||
								  statistik.KoeretoejMotorSlagVolumenIkkeTilgaengelig != "NULL" ||
								  statistik.KoeretoejMotorStandStoej != "NULL" ||
								  statistik.KoeretoejMotorStandStoejOmdrejningstal != "NULL" ||
								  statistik.KoeretoejMotorStoersteEffekt != "NULL" ||
								  statistik.KoeretoejMotorStoersteEffektIkkeTilgaengelig != "NULL" ||
								  statistik.KoeretoejMotorInnovativTeknik != "NULL" ||
								  statistik.KoeretoejMotorInnovativTeknikAntal != "NULL" ||
								  statistik.KoeretoejMotorElektriskForbrug != "NULL" ||
								  statistik.KoeretoejMotorFuelmode != "NULL" ||
								  statistik.KoeretoejMotorGasforbrug != "NULL" ||
								  statistik.KoeretoejMotorElektriskRaekkevidde != "NULL" ||
								  statistik.KoeretoejMotorBatterikapacitet != "NULL" ||
								  statistik.KoeretoejMotorBraendstofforbrugMaalt != "NULL" ||
								  statistik.KoeretoejMotorElektriskForbrugMaalt != "NULL" ||
								  statistik.KoeretoejMotorMaaleNormTypeNavn != "NULL" ||
								  statistik.KoeretoejMotorMaaleNormTypeNummer != "NULL" ||
								  statistik.KoeretoejMotorCO2UdslipBeregnet != "NULL" ||
								  statistik.KoeretoejMotorBraendselscelle != "NULL" ||
								  statistik.KoeretoejMotorDrivmiddelPrimaer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejMotorStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejMotorCylinderAntal,
									  statistik.KoeretoejMotorSlagVolumen,
									  statistik.KoeretoejMotorSlagVolumenIkkeTilgaengelig,
									  statistik.KoeretoejMotorStoersteEffekt,
									  statistik.KoeretoejMotorStoersteEffektIkkeTilgaengelig,
									  statistik.KoeretoejMotorKilometerstand,
									  statistik.KoeretoejMotorKilometerstandDokumentation,
									  statistik.KoeretoejMotorKilometerstandIkkeTilgaengelig,
									  statistik.KoeretoejMotorKmPerLiter,
									  statistik.KoeretoejMotorKMPerLiterPreCalc,
									  statistik.KoeretoejMotorPlugInHybrid,
									  statistik.KoeretoejMotorMaerkning,
									  statistik.KoeretoejMotorStandStoej,
									  statistik.KoeretoejMotorKoerselStoej,
									  statistik.KoeretoejMotorStandStoejOmdrejningstal,
									  statistik.KoeretoejMotorInnovativTeknik,
									  statistik.KoeretoejMotorInnovativTeknikAntal,
									  statistik.KoeretoejMotorElektriskForbrug,
									  statistik.KoeretoejMotorFuelmode,
									  statistik.KoeretoejMotorGasforbrug,
									  statistik.KoeretoejMotorElektriskRaekkevidde,
									  statistik.KoeretoejMotorBatterikapacitet,
									  statistik.KoeretoejMotorBraendstofforbrugMaalt,
									  statistik.KoeretoejMotorElektriskForbrugMaalt,
									  statistik.KoeretoejMotorMaaleNormTypeNavn,
									  statistik.KoeretoejMotorMaaleNormTypeNummer,
									  statistik.KoeretoejMotorCO2UdslipBeregnet,
									  statistik.KoeretoejMotorBraendselscelle,
									  statistik.KoeretoejMotorDrivmiddelPrimaer,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejMotorStatement:", err)
								  }
							  }

							  // FOR-LOOP FOR KoeretoejDrivmiddelSamling
							  for i := range statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur {
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBraendselscelle = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBraendselscelle)) //YC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorPlugInHybrid = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorPlugInHybrid)) //NC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorDrivmiddelPrimaer = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorDrivmiddelPrimaer)) //YC

								  if (statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBraendselscelle != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorPlugInHybrid != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorDrivmiddelPrimaer != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejDrivmiddelStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBraendselscelle,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorPlugInHybrid,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorDrivmiddelPrimaer,

									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejDrivmiddelStatement in for-loop:", err)
									  }
								  }

								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorMaaleNormTypeNavn = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorMaaleNormTypeNavn) //YC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorMaaleNormTypeNummer = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorMaaleNormTypeNummer) //YC

								  if (statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorMaaleNormTypeNavn != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorMaaleNormTypeNummer != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejMaaleNormStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorMaaleNormTypeNummer,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorMaaleNormTypeNavn,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejMaaleNormStatement in for-loop:", err)
									  }
								  }

								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].DrivkraftTypeNavn = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].DrivkraftTypeNavn) //YC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].DrivkraftTypeNummer = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].DrivkraftTypeNummer) //YC

								  if (statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].DrivkraftTypeNavn != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].DrivkraftTypeNummer != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejDrivkraftTypeStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].DrivkraftTypeNummer,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].DrivkraftTypeNavn,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejDrivkraftTypeStatement in for-loop:", err)
									  }
								  }

								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorKmPerLiter = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorKmPerLiter) //YC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorKMPerLiterPreCalc = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorKMPerLiterPreCalc) //NC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBraendstofforbrugMaalt = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBraendstofforbrugMaalt) //NC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorGasforbrug = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorGasforbrug) //NC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorCO2UdslipBeregnet = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorCO2UdslipBeregnet) //NC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMiljoeOplysningCO2Udslip = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMiljoeOplysningCO2Udslip) //YC

								  if (statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorKmPerLiter != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorKMPerLiterPreCalc != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBraendstofforbrugMaalt != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorGasforbrug != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorCO2UdslipBeregnet != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMiljoeOplysningCO2Udslip != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejBraendstofStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorKmPerLiter,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorKMPerLiterPreCalc,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBraendstofforbrugMaalt,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorGasforbrug,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorCO2UdslipBeregnet,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMiljoeOplysningCO2Udslip,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejBraendstofStatement in for-loop:", err)
									  }
								  }

								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskForbrug = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskForbrug) //YC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskForbrugMaalt = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskForbrugMaalt) //NC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskRaekkevidde = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskRaekkevidde) //NC
								  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBatterikapacitet = trimAndSetEmptyToNull(statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBatterikapacitet) //NC

								  if (statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskForbrug != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskForbrugMaalt != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskRaekkevidde != "NULL" ||
									  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBatterikapacitet != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejElforbrugStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskForbrug,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskForbrugMaalt,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorElektriskRaekkevidde,
										  statistik.KoeretoejDrivmiddelSamling.DrivmiddelStruktur[i].KoeretoejMotorBatterikapacitet,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejElforbrugStatement in for-loop:", err)
									  }
								  }
							  }

							  // FOR-LOOP FOR DispensationTypeSamling
							  for i := range statistik.DispensationTypeSamling.DispensationTypeStruktur {
								  statistik.DispensationTypeSamling.DispensationTypeStruktur[i].DispensationTypeNummer = trimAndSetEmptyToNull(statistik.DispensationTypeSamling.DispensationTypeStruktur[i].DispensationTypeNummer) //NC
								  statistik.DispensationTypeSamling.DispensationTypeStruktur[i].DispensationTypeNavn = trimAndSetEmptyToNull(statistik.DispensationTypeSamling.DispensationTypeStruktur[i].DispensationTypeNavn) //NC
								  statistik.DispensationTypeSamling.DispensationTypeStruktur[i].KoeretoejDispensationTypeKommentar = trimAndSetEmptyToNull(statistik.DispensationTypeSamling.DispensationTypeStruktur[i].KoeretoejDispensationTypeKommentar) //NC

								  if (statistik.DispensationTypeSamling.DispensationTypeStruktur[i].DispensationTypeNummer != "NULL" ||
									  statistik.DispensationTypeSamling.DispensationTypeStruktur[i].DispensationTypeNavn != "NULL" ||
									  statistik.DispensationTypeSamling.DispensationTypeStruktur[i].KoeretoejDispensationTypeKommentar != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejDispensationTypeStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.DispensationTypeSamling.DispensationTypeStruktur[i].DispensationTypeNummer,
										  statistik.DispensationTypeSamling.DispensationTypeStruktur[i].DispensationTypeNavn,
										  statistik.DispensationTypeSamling.DispensationTypeStruktur[i].KoeretoejDispensationTypeKommentar,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejDispensationTypeStatement in for-loop:", err)
									  }
								  }
							  }

							  // FOR-LOOP FOR KoeretoejUdstyrSamling
							  for i := range statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur {
								  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrAntal = trimAndSetEmptyToNull(statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrAntal) //YC

								  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeNavn = trimAndSetEmptyToNull(statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeNavn) //YC
								  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeNummer = trimAndSetEmptyToNull(statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeNummer) //YC
								  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeStandardAntal = trimAndSetEmptyToNull(statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeStandardAntal) //NC
								  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedForespoergsel = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedForespoergsel)) //YC
								  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedStandardOprettelse = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedStandardOprettelse)) //YC
								  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedSyn = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedSyn)) //YC

								  if (statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrAntal != "NULL" ||

									  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeNavn != "NULL" ||
									  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeNummer != "NULL" ||
									  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeStandardAntal != "NULL" ||
									  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedForespoergsel != "NULL" ||
									  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedStandardOprettelse != "NULL" ||
									  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedSyn != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejUdstyrStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrAntal,

										  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeNummer,
										  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeNavn,
										  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedSyn,
										  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedForespoergsel,
										  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeVisesVedStandardOprettelse,
										  statistik.KoeretoejUdstyrSamling.KoeretoejUdstyrStruktur[i].KoeretoejUdstyrTypeStandardAntal,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejUdstyrStatement in for-loop:", err)
									  }
								  }
							  }

							  // FOR-LOOP FOR KoeretoejBlokeringAarsagListe
							  for i := range statistik.KoeretoejBlokeringAarsagListe.KoeretoejBlokeringAarsag {
								  statistik.KoeretoejBlokeringAarsagListe.KoeretoejBlokeringAarsag[i].KoeretoejBlokeringAarsagTypeNavn = trimAndSetEmptyToNull(statistik.KoeretoejBlokeringAarsagListe.KoeretoejBlokeringAarsag[i].KoeretoejBlokeringAarsagTypeNavn) //YC
								  statistik.KoeretoejBlokeringAarsagListe.KoeretoejBlokeringAarsag[i].KoeretoejBlokeringAarsagTypeNummer = trimAndSetEmptyToNull(statistik.KoeretoejBlokeringAarsagListe.KoeretoejBlokeringAarsag[i].KoeretoejBlokeringAarsagTypeNummer) //YC

								  if (statistik.KoeretoejBlokeringAarsagListe.KoeretoejBlokeringAarsag[i].KoeretoejBlokeringAarsagTypeNavn != "NULL" ||
									  statistik.KoeretoejBlokeringAarsagListe.KoeretoejBlokeringAarsag[i].KoeretoejBlokeringAarsagTypeNummer != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejBlokeringAarsagStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.KoeretoejBlokeringAarsagListe.KoeretoejBlokeringAarsag[i].KoeretoejBlokeringAarsagTypeNummer,
										  statistik.KoeretoejBlokeringAarsagListe.KoeretoejBlokeringAarsag[i].KoeretoejBlokeringAarsagTypeNavn,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejBlokeringAarsagStatement in for-loop:", err)
									  }
								  }
							  }

							  statistik.PrisOplysningerMindsteBeskatningspris = trimAndSetEmptyToNull(statistik.PrisOplysningerMindsteBeskatningspris) //NC
							  statistik.PrisOplysningerIndkoebsPris = trimAndSetEmptyToNull(statistik.PrisOplysningerIndkoebsPris) //NC
							  statistik.PrisOplysningerStandardPris = trimAndSetEmptyToNull(statistik.PrisOplysningerStandardPris) //NC

							  if (statistik.PrisOplysningerMindsteBeskatningspris != "NULL" ||
								  statistik.PrisOplysningerIndkoebsPris != "NULL" ||
								  statistik.PrisOplysningerStandardPris != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejPrisOplysningerStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.PrisOplysningerStandardPris,
									  statistik.PrisOplysningerIndkoebsPris,
									  statistik.PrisOplysningerMindsteBeskatningspris,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejPrisOplysningerStatement:", err)
								  }
							  }

							  statistik.KoeretoejGruppeNavn = trimAndSetEmptyToNull(statistik.KoeretoejGruppeNavn) //NC
							  statistik.KoeretoejGruppeNummer = trimAndSetEmptyToNull(statistik.KoeretoejGruppeNummer) //NC

							  if (statistik.KoeretoejGruppeNavn != "NULL" ||
								  statistik.KoeretoejGruppeNummer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejGruppeStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejGruppeNummer,
									  statistik.KoeretoejGruppeNavn,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejGruppeStatement:", err)
								  }
							  }

							  statistik.KoeretoejUndergruppeNavn = trimAndSetEmptyToNull(statistik.KoeretoejUndergruppeNavn) //NC
							  statistik.KoeretoejUndergruppeNummer = trimAndSetEmptyToNull(statistik.KoeretoejUndergruppeNummer) //NC

							  if (statistik.KoeretoejUndergruppeNavn != "NULL" ||
								  statistik.KoeretoejUndergruppeNummer != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejUndergruppeStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejUndergruppeNummer,
									  statistik.KoeretoejUndergruppeNavn,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejUndergruppeStatement:", err)
								  }
							  }

							  // FUTURE WORKS: FOR-LOOP FOR EjerBrugerSamling

							  // statistik.AdresseFortloebendeNummer = trimAndSetEmptyToNull(statistik.AdresseFortloebendeNummer) //N
							  // statistik.AdresseAnvendelseKode = trimAndSetEmptyToNull(statistik.AdresseAnvendelseKode) //N
							  // statistik.AdresseVejNavn = trimAndSetEmptyToNull(statistik.AdresseVejNavn) //N
							  // statistik.AdresseVejKode = trimAndSetEmptyToNull(statistik.AdresseVejKode) //N
							  // statistik.AdresseFraHusNummer = trimAndSetEmptyToNull(statistik.AdresseFraHusNummer) //N
							  // statistik.AdresseFraHusBogstav = trimAndSetEmptyToNull(statistik.AdresseFraHusBogstav) //N
							  // statistik.AdresseTilHusNummer = trimAndSetEmptyToNull(statistik.AdresseTilHusNummer) //N
							  // statistik.AdresseTilHusBogstav = trimAndSetEmptyToNull(statistik.AdresseTilHusBogstav) //N
							  // statistik.AdresseLigeUlige = trimAndSetEmptyToNull(statistik.AdresseLigeUlige) //N
							  // statistik.AdresseLejlighedNummer = trimAndSetEmptyToNull(statistik.AdresseLejlighedNummer) //N
							  // statistik.AdresseHusNavn = trimAndSetEmptyToNull(statistik.AdresseHusNavn) //N
							  // statistik.AdresseEtage = trimAndSetEmptyToNull(statistik.AdresseEtage) //N
							  // statistik.AdresseEtageTekst = trimAndSetEmptyToNull(statistik.AdresseEtageTekst) //N
							  // statistik.AdresseSideDoerTekst = trimAndSetEmptyToNull(statistik.AdresseSideDoerTekst) //N
							  // statistik.AdresseCONavn = trimAndSetEmptyToNull(statistik.AdresseCONavn) //N
							  statistik.AdressePostNummer = trimAndSetEmptyToNull(statistik.AdressePostNummer) //NC
							  // statistik.AdressePostDistrikt = trimAndSetEmptyToNull(statistik.AdressePostDistrikt) //N
							  // statistik.AdresseLandsBy = trimAndSetEmptyToNull(statistik.AdresseLandsBy) //N
							  // statistik.AdresseByNavn = trimAndSetEmptyToNull(statistik.AdresseByNavn) //N
							  // statistik.AdresseLandsDel = trimAndSetEmptyToNull(statistik.AdresseLandsDel) //N
							  // statistik.AdressePostBox = trimAndSetEmptyToNull(statistik.AdressePostBox) //N
							  // statistik.AdresseGyldigFra = trimAndSetEmptyToNull(statistik.AdresseGyldigFra) //N
							  // statistik.AdresseGyldigTil = trimAndSetEmptyToNull(statistik.AdresseGyldigTil) //N

							  /* FUTURE WORKS:
							  if (statistik.AdresseFortloebendeNummer != "NULL" ||
								  statistik.AdresseAnvendelseKode != "NULL" ||
								  statistik.AdresseVejNavn != "NULL" ||
								  statistik.AdresseVejKode != "NULL" ||
								  statistik.AdresseFraHusNummer != "NULL" ||
								  statistik.AdresseFraHusBogstav != "NULL" ||
								  statistik.AdresseTilHusNummer != "NULL" ||
								  statistik.AdresseTilHusBogstav != "NULL" ||
								  statistik.AdresseLigeUlige != "NULL" ||
								  statistik.AdresseLejlighedNummer != "NULL" ||
								  statistik.AdresseHusNavn != "NULL" ||
								  statistik.AdresseEtage != "NULL" ||
								  statistik.AdresseEtageTekst != "NULL" ||
								  statistik.AdresseSideDoerTekst != "NULL" ||
								  statistik.AdresseCONavn != "NULL" ||
								  statistik.AdressePostNummer != "NULL" ||
								  statistik.AdressePostDistrikt != "NULL" ||
								  statistik.AdresseLandsBy != "NULL" ||
								  statistik.AdresseByNavn != "NULL" ||
								  statistik.AdresseLandsDel != "NULL" ||
								  statistik.AdressePostBox != "NULL" ||
								  statistik.AdresseGyldigFra != "NULL" ||
								  statistik.AdresseGyldigTil != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejAdresseStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.AdresseFortloebendeNummer,
									  statistik.AdresseAnvendelseKode,
									  statistik.AdresseVejNavn,
									  statistik.AdresseVejKode,
									  statistik.AdresseFraHusNummer,
									  statistik.AdresseFraHusBogstav,
									  statistik.AdresseTilHusNummer,
									  statistik.AdresseTilHusBogstav,
									  statistik.AdresseLigeUlige,
									  statistik.AdresseLejlighedNummer,
									  statistik.AdresseHusNavn,
									  statistik.AdresseEtage,
									  statistik.AdresseEtageTekst,
									  statistik.AdresseSideDoerTekst,
									  statistik.AdresseCONavn,
									  statistik.AdressePostNummer,
									  statistik.AdressePostDistrikt,
									  statistik.AdresseLandsBy,
									  statistik.AdresseByNavn,
									  statistik.AdresseLandsDel,
									  statistik.AdressePostBox,
									  statistik.AdresseGyldigFra,
									  statistik.AdresseGyldigTil,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejAdresseStatement:", err)
								  }
							  }
							  */

							  statistik.SynResultatNummer = trimAndSetEmptyToNull(statistik.SynResultatNummer) //M
							  statistik.SynResultatSynsDato = trimAndSetEmptyToNull(statistik.SynResultatSynsDato) //YC
							  statistik.SynResultatSynsResultat = trimAndSetEmptyToNull(statistik.SynResultatSynsResultat) //YC
							  statistik.SynResultatSynStatus = trimAndSetEmptyToNull(statistik.SynResultatSynStatus) //YC
							  statistik.SynResultatSynStatusDato = trimAndSetEmptyToNull(statistik.SynResultatSynStatusDato) //YC
							  statistik.SynResultatSynsType = trimAndSetEmptyToNull(statistik.SynResultatSynsType) //YC
							  statistik.SynResultatOmsynMoedeDato = trimAndSetEmptyToNull(statistik.SynResultatOmsynMoedeDato) //M
							  statistik.SynResultatKoeretoejMotorKilometerstand = trimAndSetEmptyToNull(statistik.SynResultatKoeretoejMotorKilometerstand) //YC

							  if (statistik.SynResultatNummer != "NULL" ||
								  statistik.SynResultatSynsDato != "NULL" ||
								  statistik.SynResultatSynsResultat != "NULL" ||
								  statistik.SynResultatSynStatus != "NULL" ||
								  statistik.SynResultatSynStatusDato != "NULL" ||
								  statistik.SynResultatSynsType != "NULL" ||
								  statistik.SynResultatOmsynMoedeDato != "NULL" ||
								  statistik.SynResultatKoeretoejMotorKilometerstand != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejSynResultatStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.SynResultatNummer,
									  statistik.SynResultatKoeretoejMotorKilometerstand,
									  statistik.SynResultatSynStatus,
									  statistik.SynResultatSynStatusDato,
									  statistik.SynResultatSynsDato,
									  statistik.SynResultatSynsResultat,
									  statistik.SynResultatSynsType,
									  statistik.SynResultatOmsynMoedeDato,
								  )
								  if err != nil {
									fmt.Println("Error executing SQL koeretoejSynResultatStatement:", err)
								  }
							  }

							  // FUTURE WORKS: FOR-LOOP FOR KoeretoejRegistreringGrundlagSamlingStrukturType

							  statistik.KoeretoejRegistreringGyldigFra = trimAndSetEmptyToNull(statistik.KoeretoejRegistreringGyldigFra) //NC
							  statistik.KoeretoejRegistreringGyldigTil = trimAndSetEmptyToNull(statistik.KoeretoejRegistreringGyldigTil) //NC
							  statistik.KoeretoejRegistreringNummer = trimAndSetEmptyToNull(statistik.KoeretoejRegistreringNummer) //M
							  statistik.KoeretoejRegistreringStatus = trimAndSetEmptyToNull(statistik.KoeretoejRegistreringStatus) //YC
							  statistik.KoeretoejRegistreringStatusDato = trimAndSetEmptyToNull(statistik.KoeretoejRegistreringStatusDato) //YC
							  statistik.KoeretoejRegistreringStatusAarsag = trimAndSetEmptyToNull(statistik.KoeretoejRegistreringStatusAarsag) //NC
							  statistik.KoeretoejRegistreringKontrolTal = trimAndSetEmptyToNull(statistik.KoeretoejRegistreringKontrolTal) //NC
							  statistik.KoeretoejRegistreringGrundlagIdent = trimAndSetEmptyToNull(statistik.KoeretoejRegistreringGrundlagIdent) //NC
							  statistik.KoeretoejRegistreringSenesteHaendelse = trimAndSetEmptyToNull(statistik.KoeretoejRegistreringSenesteHaendelse) //NC
							  statistik.KoeretoejRegistreringTilknyttetLeasingForhold = matchTrueFalse(trimAndSetEmptyToNull(statistik.KoeretoejRegistreringTilknyttetLeasingForhold)) //NC

							  if (statistik.KoeretoejRegistreringGyldigFra != "NULL" ||
								  statistik.KoeretoejRegistreringGyldigTil != "NULL" ||
								  statistik.KoeretoejRegistreringNummer != "NULL" ||
								  statistik.KoeretoejRegistreringStatus != "NULL" ||
								  statistik.KoeretoejRegistreringStatusDato != "NULL" ||
								  statistik.KoeretoejRegistreringStatusAarsag != "NULL" ||
								  statistik.KoeretoejRegistreringKontrolTal != "NULL" ||
								  statistik.KoeretoejRegistreringGrundlagIdent != "NULL" ||
								  statistik.KoeretoejRegistreringSenesteHaendelse != "NULL" ||
								  statistik.KoeretoejRegistreringTilknyttetLeasingForhold != "NULL") {

								_, err = tx.StmtContext(ctx, koeretoejRegistreringStatement).
									ExecContext(ctx,
									  statistik.KoeretoejIdent,

									  statistik.KoeretoejRegistreringNummer,
									  statistik.KoeretoejRegistreringStatus,
									  statistik.KoeretoejRegistreringStatusDato,
									  statistik.KoeretoejRegistreringStatusAarsag,
									  statistik.KoeretoejRegistreringKontrolTal,
									  statistik.KoeretoejRegistreringGyldigFra,
									  statistik.KoeretoejRegistreringGyldigTil,
									  statistik.KoeretoejRegistreringGrundlagIdent,
									  statistik.KoeretoejRegistreringSenesteHaendelse,
									  statistik.KoeretoejRegistreringTilknyttetLeasingForhold,
								  )
								  if err != nil {
									  fmt.Println("Error executing SQL koeretoejRegistreringStatement:", err)
								  }
							  }

							  // FOR-LOOP FOR Tilladelse
							  for i := range statistik.Tilladelse.TilladelseStruktur {
								  statistik.Tilladelse.TilladelseStruktur[i].TilladelseGyldigFra = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].TilladelseGyldigFra) //YC
								  statistik.Tilladelse.TilladelseStruktur[i].TilladelseGyldigTil = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].TilladelseGyldigTil) //NC
								  statistik.Tilladelse.TilladelseStruktur[i].TilladelseKommentar = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].TilladelseKommentar) //YC
								  statistik.Tilladelse.TilladelseStruktur[i].TilladelseNummer = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].TilladelseNummer) //M
								  statistik.Tilladelse.TilladelseStruktur[i].TilladelseKunGodkendtForRegistreretEjer = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].TilladelseKunGodkendtForRegistreretEjer) //M
								  statistik.Tilladelse.TilladelseStruktur[i].TilladelseKombinationKoeretoejIdent = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].TilladelseKombinationKoeretoejIdent) //M

								  if (statistik.Tilladelse.TilladelseStruktur[i].TilladelseGyldigFra != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].TilladelseGyldigTil != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].TilladelseKommentar != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].TilladelseNummer != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].TilladelseKunGodkendtForRegistreretEjer != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].TilladelseKombinationKoeretoejIdent != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejTilladelseStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.Tilladelse.TilladelseStruktur[i].TilladelseNummer,
										  statistik.Tilladelse.TilladelseStruktur[i].TilladelseGyldigFra,
										  statistik.Tilladelse.TilladelseStruktur[i].TilladelseGyldigTil,
										  statistik.Tilladelse.TilladelseStruktur[i].TilladelseKommentar,
										  statistik.Tilladelse.TilladelseStruktur[i].TilladelseKunGodkendtForRegistreretEjer,
										  statistik.Tilladelse.TilladelseStruktur[i].TilladelseKombinationKoeretoejIdent,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejTilladelseStatement in for-loop:", err)
									  }
								  }

								  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeNavn = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeNavn) //YC
								  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeNummer = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeNummer) //YC
								  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeErPeriodeBegraenset = matchTrueFalse(trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeErPeriodeBegraenset)) //M
								  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypePeriodeLaengde = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypePeriodeLaengde) //M

								  if (statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeNavn != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeNummer != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeErPeriodeBegraenset != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypePeriodeLaengde != "NULL") {

									_, err = tx.StmtContext(ctx, koeretoejTilladelseTypeStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeNummer,
										  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeNavn,
										  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypeErPeriodeBegraenset,
										  statistik.Tilladelse.TilladelseStruktur[i].TilladelseTypePeriodeLaengde,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejTilladelseTypeStatement in for-loop:", err)
									  }
								  }

								  statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejIdent = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejIdent) //YC
								  statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationRegistreringNummerNummer = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationRegistreringNummerNummer) //YC
								  statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejOplysningStelNummer = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejOplysningStelNummer) //NC

								  if (statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejIdent != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationRegistreringNummerNummer != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejOplysningStelNummer != "NULL") {

									  if (statistik.KoeretoejIdent != statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejIdent) {
										  // fmt.Printf(
										  // 	"\nstatistik.KoeretoejIdent:\n%s\ndoes not equal statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejIdent:\n%s\n",
										  // 	statistik.KoeretoejIdent,
										  // 	statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejIdent,
										  // )
									  }

									_, err = tx.StmtContext(ctx, koeretoejTilladelseTypeVariabelKombinationStatement).
										ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejIdent,
										  statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationRegistreringNummerNummer,
										  statistik.Tilladelse.TilladelseStruktur[i].VariabelKombinationKoeretoejOplysningStelNummer,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejTilladelseTypeVariabelKombinationStatement in for-loop:", err)
									  }
								  }

								  statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejIdent = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejIdent) //YC
								  statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejOplysningStelNummer = trimAndSetEmptyToNull(statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejOplysningStelNummer) //NC

								  if (statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejIdent != "NULL" ||
									  statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejOplysningStelNummer != "NULL") {

									  if (statistik.KoeretoejIdent != statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejIdent) {
										  // fmt.Printf(
										  // 	"\nstatistik.KoeretoejIdent:\n%s\ndoes not equal statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejIdent:\n%s\n",
										  // 	statistik.KoeretoejIdent,
										  // 	statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejIdent,
										  // )
									  }

									_, err = tx.StmtContext(ctx, koeretoejTilladelseTypeFastTilkoblingStatement).
									ExecContext(ctx,
										  statistik.KoeretoejIdent,

										  statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejIdent,
										  statistik.Tilladelse.TilladelseStruktur[i].FastTilkoblingKoeretoejOplysningStelNummer,
									  )
									  if err != nil {
										  fmt.Println("Error executing SQL koeretoejTilladelseTypeFastTilkoblingStatement in for-loop:", err)
									  }
								  }
							  }
						  }
				  // … repeat for the other statements, or factor into a helper …

				  // 6) batching logic
				  inBatch++
				  if inBatch >= batchSize {
					  if err := tx.Commit(); err != nil {
						  log.Printf("[worker %d] commit: %v", workerID, err)
						  cancel()
						  return
					  }
					  // start a fresh tx
					  tx, err = db.BeginTx(ctx, nil)
					  if err != nil {
						  log.Printf("[worker %d] new tx: %v", workerID, err)
						  cancel()
						  return
					  }
					  inBatch = 0
					  cnt := atomic.AddInt64(&total, 1)
		              if cnt%batchSize == 0 {
					      fmt.Printf("\rRow inserted (%d/%d)", cnt, 13217216)
					  }
				  }
			  }
		  }
	  }(i)
  }

	for {
	  tok, err := decoder.Token()
	  if err != nil {
		  if err == io.EOF {
			  break
		  }
		  log.Fatal(err)
	  }
	  if se, ok := tok.(xml.StartElement); ok && se.Name.Local == "Statistik" {
		  var s Statistik
		  if err := decoder.DecodeElement(&s, &se); err != nil {
			  log.Fatal(err)
		  }
		  select {
		  case statistikCh <- s:
		  case <-ctx.Done():
			  log.Fatal("aborting due to worker error")
		  }
	  }
  }

  close(statistikCh)
  wg.Wait()

  fmt.Println("All done!  Elapsed:", time.Since(startTime))
}

func trimAndSetEmptyToNull(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "NULL"
	}
	return s
}

func matchTrueFalse(s string) string {
	switch strings.ToLower(s) {
	case "true":
		return "1"
	case "false":
		return "0"
	default:
		return s
	}
}
