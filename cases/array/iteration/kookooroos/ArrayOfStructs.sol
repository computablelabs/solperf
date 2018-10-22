pragma solidity 0.4.25;

/**
  Use case: We need to know if a given value, for a given key, is present within any struct that is contained within an array.

  Say we have an array that holds structs. What the structs hold is largely irrevalent here,
  (though there may be a case where we want to test the if the size of the struct matters)
  what we are interested in is an "owner" attribute which holds an Adress. The presence of a given adress in any
  struct within this array means that the given address is an "owner" of whatever this struct holds.

  For this example we'll cite that magnum-opus dissertation of socioeconomic status meets price-of-fame meets fleeting-nature-of-success masterwork:
  Get Him to the Greek:
    "I own 20 Koo Koo Roos."

                "21 sir..."

    "I own 21 Koo Koo Roos. Y'all don't own one Koo Koo Roo. Not one..."


  NOTE: A second constraint is important here. It is a *known law* that no one may own more than 20 Koo Koo Roos unless they are
  Sean Combs. In a solution with only an array of structs, we'll have to count each every time...
*/

contract ArrayOfStructs {
  // We will assume diddy is not part of our test.
  uint public MAX_KOO_KOO_ROOS = 20; // TODO - not implemented upper bound in this case yet

  struct KooKooRoo {
    address owner;
    // none of these others really matter, but in any actual struct there'll be other stuff...
    uint listed; // maybe a timestamp of when we listed the location
    bytes32 location; // where it was. We'll hash this for the mapping key
  }

  KooKooRoo[] public kooKooRoos;

  /**
  @dev Method to create an owned koo koo roo in the kookooroos array. NOTE this does not prevent duplicates, that will require another
  data structure, or us to iterate the current array each time and count... TODO
  NOTE: The msg.sender must not exceed the MAX_KOO_KOO_ROOS upper bound. TODO
  */
  function becomeOwner(string location) external returns (bool) {
    // create the ownable koo koo roo, use the converted location as the key
    bytes32 converted = stringToBytes32(location);
    uint n = kooKooRoos.push(KooKooRoo(msg.sender, now, converted));
    return n > 0;
  }

  function isKooKooRooOwner(address candidate) external view returns (bool) {
    // TODO
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
