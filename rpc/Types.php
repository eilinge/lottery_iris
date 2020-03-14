<?php
namespace rpc;

/**
 * Autogenerated by Thrift Compiler (0.11.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
use Thrift\Base\TBase;
use Thrift\Type\TType;
use Thrift\Type\TMessageType;
use Thrift\Exception\TException;
use Thrift\Exception\TProtocolException;
use Thrift\Protocol\TProtocol;
use Thrift\Protocol\TBinaryProtocolAccelerated;
use Thrift\Exception\TApplicationException;


class DataGiftPrize {
  static $isValidate = false;

  static $_TSPEC = array(
    1 => array(
      'var' => 'Id',
      'isRequired' => false,
      'type' => TType::I64,
      ),
    2 => array(
      'var' => 'Title',
      'isRequired' => false,
      'type' => TType::STRING,
      ),
    3 => array(
      'var' => 'Img',
      'isRequired' => false,
      'type' => TType::STRING,
      ),
    4 => array(
      'var' => 'Displayorder',
      'isRequired' => false,
      'type' => TType::I64,
      ),
    5 => array(
      'var' => 'Gtype',
      'isRequired' => false,
      'type' => TType::I64,
      ),
    6 => array(
      'var' => 'Gdata',
      'isRequired' => false,
      'type' => TType::STRING,
      ),
    );

  /**
   * @var int
   */
  public $Id = 0;
  /**
   * @var string
   */
  public $Title = "";
  /**
   * @var string
   */
  public $Img = "";
  /**
   * @var int
   */
  public $Displayorder = 0;
  /**
   * @var int
   */
  public $Gtype = 0;
  /**
   * @var string
   */
  public $Gdata = "";

  public function __construct($vals=null) {
    if (is_array($vals)) {
      if (isset($vals['Id'])) {
        $this->Id = $vals['Id'];
      }
      if (isset($vals['Title'])) {
        $this->Title = $vals['Title'];
      }
      if (isset($vals['Img'])) {
        $this->Img = $vals['Img'];
      }
      if (isset($vals['Displayorder'])) {
        $this->Displayorder = $vals['Displayorder'];
      }
      if (isset($vals['Gtype'])) {
        $this->Gtype = $vals['Gtype'];
      }
      if (isset($vals['Gdata'])) {
        $this->Gdata = $vals['Gdata'];
      }
    }
  }

  public function getName() {
    return 'DataGiftPrize';
  }

  public function read($input)
  {
    $xfer = 0;
    $fname = null;
    $ftype = 0;
    $fid = 0;
    $xfer += $input->readStructBegin($fname);
    while (true)
    {
      $xfer += $input->readFieldBegin($fname, $ftype, $fid);
      if ($ftype == TType::STOP) {
        break;
      }
      switch ($fid)
      {
        case 1:
          if ($ftype == TType::I64) {
            $xfer += $input->readI64($this->Id);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        case 2:
          if ($ftype == TType::STRING) {
            $xfer += $input->readString($this->Title);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        case 3:
          if ($ftype == TType::STRING) {
            $xfer += $input->readString($this->Img);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        case 4:
          if ($ftype == TType::I64) {
            $xfer += $input->readI64($this->Displayorder);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        case 5:
          if ($ftype == TType::I64) {
            $xfer += $input->readI64($this->Gtype);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        case 6:
          if ($ftype == TType::STRING) {
            $xfer += $input->readString($this->Gdata);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        default:
          $xfer += $input->skip($ftype);
          break;
      }
      $xfer += $input->readFieldEnd();
    }
    $xfer += $input->readStructEnd();
    return $xfer;
  }

  public function write($output) {
    $xfer = 0;
    $xfer += $output->writeStructBegin('DataGiftPrize');
    if ($this->Id !== null) {
      $xfer += $output->writeFieldBegin('Id', TType::I64, 1);
      $xfer += $output->writeI64($this->Id);
      $xfer += $output->writeFieldEnd();
    }
    if ($this->Title !== null) {
      $xfer += $output->writeFieldBegin('Title', TType::STRING, 2);
      $xfer += $output->writeString($this->Title);
      $xfer += $output->writeFieldEnd();
    }
    if ($this->Img !== null) {
      $xfer += $output->writeFieldBegin('Img', TType::STRING, 3);
      $xfer += $output->writeString($this->Img);
      $xfer += $output->writeFieldEnd();
    }
    if ($this->Displayorder !== null) {
      $xfer += $output->writeFieldBegin('Displayorder', TType::I64, 4);
      $xfer += $output->writeI64($this->Displayorder);
      $xfer += $output->writeFieldEnd();
    }
    if ($this->Gtype !== null) {
      $xfer += $output->writeFieldBegin('Gtype', TType::I64, 5);
      $xfer += $output->writeI64($this->Gtype);
      $xfer += $output->writeFieldEnd();
    }
    if ($this->Gdata !== null) {
      $xfer += $output->writeFieldBegin('Gdata', TType::STRING, 6);
      $xfer += $output->writeString($this->Gdata);
      $xfer += $output->writeFieldEnd();
    }
    $xfer += $output->writeFieldStop();
    $xfer += $output->writeStructEnd();
    return $xfer;
  }

}

class DataResult {
  static $isValidate = false;

  static $_TSPEC = array(
    1 => array(
      'var' => 'Code',
      'isRequired' => false,
      'type' => TType::I64,
      ),
    2 => array(
      'var' => 'Msg',
      'isRequired' => false,
      'type' => TType::STRING,
      ),
    3 => array(
      'var' => 'Gift',
      'isRequired' => false,
      'type' => TType::STRUCT,
      'class' => '\rpc\DataGiftPrize',
      ),
    );

  /**
   * @var int
   */
  public $Code = null;
  /**
   * @var string
   */
  public $Msg = null;
  /**
   * @var \rpc\DataGiftPrize
   */
  public $Gift = null;

  public function __construct($vals=null) {
    if (is_array($vals)) {
      if (isset($vals['Code'])) {
        $this->Code = $vals['Code'];
      }
      if (isset($vals['Msg'])) {
        $this->Msg = $vals['Msg'];
      }
      if (isset($vals['Gift'])) {
        $this->Gift = $vals['Gift'];
      }
    }
  }

  public function getName() {
    return 'DataResult';
  }

  public function read($input)
  {
    $xfer = 0;
    $fname = null;
    $ftype = 0;
    $fid = 0;
    $xfer += $input->readStructBegin($fname);
    while (true)
    {
      $xfer += $input->readFieldBegin($fname, $ftype, $fid);
      if ($ftype == TType::STOP) {
        break;
      }
      switch ($fid)
      {
        case 1:
          if ($ftype == TType::I64) {
            $xfer += $input->readI64($this->Code);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        case 2:
          if ($ftype == TType::STRING) {
            $xfer += $input->readString($this->Msg);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        case 3:
          if ($ftype == TType::STRUCT) {
            $this->Gift = new \rpc\DataGiftPrize();
            $xfer += $this->Gift->read($input);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        default:
          $xfer += $input->skip($ftype);
          break;
      }
      $xfer += $input->readFieldEnd();
    }
    $xfer += $input->readStructEnd();
    return $xfer;
  }

  public function write($output) {
    $xfer = 0;
    $xfer += $output->writeStructBegin('DataResult');
    if ($this->Code !== null) {
      $xfer += $output->writeFieldBegin('Code', TType::I64, 1);
      $xfer += $output->writeI64($this->Code);
      $xfer += $output->writeFieldEnd();
    }
    if ($this->Msg !== null) {
      $xfer += $output->writeFieldBegin('Msg', TType::STRING, 2);
      $xfer += $output->writeString($this->Msg);
      $xfer += $output->writeFieldEnd();
    }
    if ($this->Gift !== null) {
      if (!is_object($this->Gift)) {
        throw new TProtocolException('Bad type in structure.', TProtocolException::INVALID_DATA);
      }
      $xfer += $output->writeFieldBegin('Gift', TType::STRUCT, 3);
      $xfer += $this->Gift->write($output);
      $xfer += $output->writeFieldEnd();
    }
    $xfer += $output->writeFieldStop();
    $xfer += $output->writeStructEnd();
    return $xfer;
  }

}


