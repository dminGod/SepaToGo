# SepaToGo

All ISO 20022 message definition XSD files converted to go structs using [xsdgen][2].

This can be used to parse XML files that implement the corresponding XSD.

You may need to slightly modify the go code for it to work with your code - Example of using pain_001.


    // Define the wrapper struct with the first element
    type PaIn001 struct {
        XMLName          xml.Name `xml:"Document"`
        CstmrCdtTrfInitn CustomerCreditTransferInitiationV10 ` xml:"CstmrCdtTrfInitn"`
    }
    
    func ParseXML(){    
    
    	var fcBy []byte
    	var err error
    
    	// Open the file and fetch the contents:
    	fcBy, err = ioutil.ReadFile("parsers/examples/pain_001.xml")
    	if err != nil {
    		fmt.Printf("There was an error reading the file : %v", err.Error())
    		return
    	}
    
    	var doc PaIn001
    
    	// Get the XML object:
    	err = xml.Unmarshal(fcBy, &doc)
    	if err != nil {
    		fmt.Printf("There was an error in parsing the XML : %v", err.Error())
    		return
    	}
    	
    	// JSON format:
        js, _ := json.MarshalIndent(doc, "", "  ")
        fmt.Printf("%+v", string(js))
    	
    }
    
    


List of formats available :

- acmt.007.001.03
- acmt.008.001.03
- acmt.009.001.03
- acmt.010.001.03
- acmt.011.001.03
- acmt.012.001.03
- acmt.013.001.03
- acmt.014.001.03
- acmt.015.001.03
- acmt.016.001.03
- acmt.017.001.03.0
- acmt.018.001.03
- acmt.019.001.03
- acmt.020.001.03
- acmt.021.001.03
- acmt.022.001.02
- acmt.023.001.02
- acmt.024.001.02
- acmt.027.001.03
- acmt.028.001.03
- acmt.029.001.03
- acmt.030.001.02
- acmt.031.001.03
- acmt.032.001.03
- acmt.033.001.02
- acmt.034.001.03
- acmt.035.001.02
- acmt.036.001.01
- acmt.037.001.02
- auth.001.001.01
- auth.002.001.01
- auth.003.001.01
- auth.018.001.02
- auth.019.001.02
- auth.020.001.02
- auth.021.001.02
- auth.022.001.02
- auth.023.001.02
- auth.024.001.02
- auth.025.001.02
- auth.026.001.02
- auth.027.001.02
- camt.003.001.07
- camt.004.001.08
- camt.005.001.08
- camt.006.001.08
- camt.007.001.08
- camt.008.001.08
- camt.009.001.07
- camt.010.001.08
- camt.011.001.07
- camt.012.001.07
- camt.013.001.04
- camt.014.001.04
- camt.015.001.04
- camt.016.001.04
- camt.017.001.04
- camt.018.001.05
- camt.019.001.07
- camt.020.001.04
- camt.021.001.06
- camt.023.001.07
- camt.024.001.06
- camt.025.001.05
- camt.026.001.08
- camt.027.001.08
- camt.028.001.10
- camt.029.001.10
- camt.030.001.05
- camt.031.001.06
- camt.032.001.04
- camt.033.001.06
- camt.034.001.06
- camt.035.001.05
- camt.036.001.05
- camt.037.001.08
- camt.038.001.04
- camt.039.001.05
- camt.046.001.05
- camt.047.001.06
- camt.048.001.05
- camt.049.001.05
- camt.050.001.05
- camt.051.001.05
- camt.052.001.08
- camt.053.001.08
- camt.054.001.08
- camt.055.001.09
- camt.056.001.09.0
- camt.057.001.06
- camt.058.001.06
- camt.059.001.06
- camt.060.001.05
- camt.069.001.03
- camt.070.001.04
- camt.071.001.03
- camt.086.001.03
- camt.087.001.07
- camt.101.001.01
- camt.102.001.01
- camt.103.001.01
- camt.104.001.01
- pacs.002.001.11
- pacs.003.001.08
- pacs.004.001.10
- pacs.007.001.10
- pacs.008.001.09
- pacs.009.001.09
- pacs.010.001.04
- pacs.028.001.04
- pain.001.001.10.1
- pain.002.001.11
- pain.007.001.10
- pain.008.001.09
- pain.009.001.05
- pain.010.001.05
- pain.011.001.05
- pain.012.001.05
- pain.013.001.08
- pain.014.001.08
- pain.017.001.01
- pain.018.001.01
- remt.001.001.04
- remt.002.001.02



[ISO 20022 Reference][1]

[1]: https://www.iso20022.org/iso-20022-message-definitions?business-domain=1
[2]: https://godoc.org/aqwari.net/xml/xsdgen