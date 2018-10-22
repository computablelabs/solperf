pragma solidity 0.4.25;

/**
  Use case: We need to know if a given value, for a given key, is present within any struct that is contained within a mapping.

  Say we have a mapping that holds structs. What the structs hold is largely irrevalent here,
  (though there may be a case where we want to test the if the size of the struct matters)
  what we are interested in is an "owner" attribute which holds an Adress. The presence of a given adress in any
  struct within this mapping means that the given address is an "owner" of whatever this mapping holds.

  For this example we'll cite that magnum-opus dissertation of socioeconomic status meets price-of-fame meets fleeting-nature-of-success masterwork:
  Get Him to the Greek:
    "I own 20 Koo Koo Roos."

                "21 sir..."

    "I own 21 Koo Koo Roos. Y'all don't own one Koo Koo Roo. Not one..."


  NOTE: A second constraint is important here. It is a *known law* that no one may own more than 20 Koo Koo Roos unless they are
  Sean Combs. In this solution we enforce the upper bound by simply using an integer as the value for the "owner" mapping.
*/

contract MappingPlusMapping {
  // We will assume diddy is not part of our test.
  uint public MAX_KOO_KOO_ROOS = 20;

  struct KooKooRoo {
    address owner;
    // none of these others really matter, but in any actual struct there'll be other stuff...
    uint listed; // maybe a timestamp of when we listed the location
    string generalManager; // some random field that we can set separately if we want, maybe check updating with...
  }

  mapping(bytes32 => KooKooRoo) public kooKooRoos;
  // in this example this is our container for quickly asserting if any given address is a kookooroo owner, whilst
  // simulataneously enforcing the upper bound
  mapping(address => uint) public kooKooRooOwners;

  /**
  @dev Method to create an owned koo koo roo in the kookooroos mapping if the location is not already there.
  The address of the msg.sender will be set owner and we will hash the location to use as a key (you can fetch a koo koo roo by that hashed location too).
  NOTE: The msg.sender must not exceed the MAX_KOO_KOO_ROOS upper bound.
  NOTE: We will increment the number of koo koo roos any owner owns in this method
  */
  function becomeOwner(string location) external upperBounded returns (bool) {
    // if not in violation, via modifier, add as an owner (using the fact that a non existant owner entry will be 0)
    kooKooRooOwners[msg.sender] = kooKooRooOwners[msg.sender] + 1;
    // create the ownable koo koo roo, use the converted location as the key
    // NOTE: as "already known", string literals cost much more to store, TODO could be a separate test?
    bytes32 converted = stringToBytes32(location);
    KooKooRoo storage kookooroo = kooKooRoos[converted];
    // set the misc data
    kookooroo.owner = msg.sender;
    kookooroo.listed = now;
    return true;
  }

  modifier upperBounded() {
    require(kooKooRooOwners[msg.sender] < MAX_KOO_KOO_ROOS, "Error: Only Diddy may own 21 Koo Koo Roos");
    _;
  }

  function isKooKooRooOwner(address candidate) external view returns (bool) {
    return kooKooRooOwners[candidate] > 0;
  }

  // public as we will allow specs to create the same mapping key by calling it
  function stringToBytes32(string src) public pure returns (bytes32 result) {
    // for this example we'll enforce that the location string can't be longer than 32 chars, obv in a prod env you'd just shorten the string here...
    require(bytes(src).length <= 32, "Error: Location must be 32 characters or less");
    // sorry, solium, there's just no other way to do this atm
    assembly {
      result := mload(add(src, 32))
    }
  }
}
