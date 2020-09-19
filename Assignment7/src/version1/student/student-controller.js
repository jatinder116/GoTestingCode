
'use strict';

const FabricCAServices = require('fabric-ca-client');
const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');

class Student {

  async enrollAdminCertificate() {
    try {
      // load the network configuration
      const ccpPath = path.resolve('..', '..', 'gsi-network', 'Assignment7', 'connection.json');
      const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

      // Create a new CA client for interacting with the CA.
      const caInfo = ccp.certificateAuthorities['ca.gsi.com'];
      const caTLSCACerts = caInfo.tlsCACerts.pem;
      const ca = new FabricCAServices(caInfo.url, { trustedRoots: caTLSCACerts, verify: false }, caInfo.caName);

      // Create a new file system based wallet for managing identities.
      const walletPath = path.join(process.cwd(), 'wallet');

      const wallet = await Wallets.newFileSystemWallet(walletPath);

      // Check to see if we've already enrolled the admin user.
      const identity = await wallet.get('admin');
      if (identity) {
        throw { status: 2, message: global.messages.enrollAdminCertificateAlready };
      }

      // Enroll the admin user, and import the new identity into the wallet.
      const enrollment = await ca.enroll({ enrollmentID: 'admin', enrollmentSecret: 'adminpw' });
      const x509Identity = {
        credentials: {
          certificate: enrollment.certificate,
          privateKey: enrollment.key.toBytes(),
        },
        mspId: 'GSIMSP',
        type: 'X.509',
      };
      await wallet.put('admin', x509Identity);

      return Promise.resolve({ message: global.messages.enrollAdmin, status: 1 });
    } catch (err) { // catch errors
      return Promise.reject({ message: `Failed to enroll admin user "admin": ${err}`, status: 0 });
    }
  }


  async registerUserCertificate() {
    try {
      // load the network configuration
      const ccpPath = path.resolve('..', '..', 'gsi-network', 'Assignment7', 'connection.json');
      const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

      // Create a new CA client for interacting with the CA.
      const caURL = ccp.certificateAuthorities['ca.gsi.com'].url;
      const ca = new FabricCAServices(caURL);

      // Create a new file system based wallet for managing identities.
      const walletPath = path.join(process.cwd(), 'wallet');
      const wallet = await Wallets.newFileSystemWallet(walletPath);

      // Check to see if we've already enrolled the user.
      const userIdentity = await wallet.get('appstudent');
      if (userIdentity) {
        throw { status: 2, message: global.messages.registerUserCertificateAlready };
      }

      // Check to see if we've already enrolled the admin user.
      const adminIdentity = await wallet.get('admin');
      if (!adminIdentity) {
        throw { status: 2, message: global.messages.adminIdentity };
      }

      // build a user object for authenticating with the CA
      const provider = wallet.getProviderRegistry().getProvider(adminIdentity.type);
      const adminUser = await provider.getUserContext(adminIdentity, 'admin');

      // Register the user, enroll the user, and import the new identity into the wallet.
      const secret = await ca.register({
        affiliation: 'org1.department1',
        enrollmentID: 'appstudent',
        role: 'client'
      }, adminUser);
      const enrollment = await ca.enroll({
        enrollmentID: 'appstudent',
        enrollmentSecret: secret
      });
      const x509Identity = {
        credentials: {
          certificate: enrollment.certificate,
          privateKey: enrollment.key.toBytes(),
        },
        mspId: 'GSIMSP',
        type: 'X.509',
      };
      await wallet.put('appstudent', x509Identity);
      return Promise.resolve({ message: global.messages.registerUserCertificate, status: 1 });
    } catch (err) { // catch errors
      return Promise.reject({ message: `Failed to enroll admin user "admin": ${err}`, status: 0 });
    }
  }


  // student
  async studentData(body, type) {
    try {
      // load the network configuration
      const ccpPath = path.resolve('..', '..', 'gsi-network', 'Assignment7', 'connection.json');
      let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

      // Create a new file system based wallet for managing identities.
      const walletPath = path.join(process.cwd(), 'wallet');
      const wallet = await Wallets.newFileSystemWallet(walletPath);

      // Check to see if we've already enrolled the user.
      const identity = await wallet.get('appstudent');
      if (!identity) {
        throw { status: 2, message: global.messages.registerUserIdentity };
      }

      // Create a new gateway for connecting to our peer node.
      const gateway = new Gateway();
      await gateway.connect(ccp, { wallet, identity: 'appstudent', discovery: { enabled: false, asLocalhost: true } });

      // Get the network (channel) our contract is deployed to.
      const network = await gateway.getNetwork('mychannel');

      // Get the contract from the network.
      const contract = network.getContract('korecontract');

      //========== Submit/Evaluate Transactions ========================================
      let displayMessage, data = "";
      if (type == 1) {
        // Submit the add student transaction.
        await contract.submitTransaction('CreateStu', body.id, body.name, body.gender, body.city);
        let message = global.messages.addStudent;
        message = message.replace('@name@', body.name);
        displayMessage = message;
      } else if (type == 2) {
        // Evaluate the get student transaction.
        const result = await contract.evaluateTransaction('GetStu', body.id);
        let studentData = JSON.parse(result.toString());
        displayMessage = global.messages.getStudent;
        data = studentData;
      } else if (type == 3) {
        // Evaluate the Get students transaction.
        const result = await contract.evaluateTransaction('GetAllStu');
        let studentData = JSON.parse(result.toString());
        displayMessage = global.messages.getAllStudent;
        data = studentData;
      } else if (type == 4) {
        // Submit the update student transaction.
        await contract.submitTransaction('UpdateStu', body.id, body.name, body.gender, body.city);
        let message = global.messages.updateStudent;
        message = message.replace('@name@', body.name);
        displayMessage = message;
      } else {
        // Submit the delete student transaction.
        await contract.submitTransaction('DeleteStu', body.id);
        displayMessage = global.messages.getStudent;
      }
      // Disconnect from the gateway.
      await gateway.disconnect();
      return Promise.resolve(data.length > 0 ? { message: displayMessage, status: 1, data: data } : { message: displayMessage, status: 1 });
    } catch (err) { // catch errors
      return Promise.reject({ message: err.message, status: 0 });
    }
  }
}


module.exports = Student;
