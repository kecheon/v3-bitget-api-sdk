<?php


namespace bitget\model\mix\account;


class SetMarginModeReq
{
    /**
     * Deposit currency
     */
    var string $marginCoin;
    /**
     * Currency pair
     */
    var string $symbol;
    /**
     * Margin mode
     */
    var string $marginMode;

    /**
     * @return mixed
     */
    public function getMarginCoin()
    {
        return $this->marginCoin;
    }

    /**
     * @param mixed $marginCoin
     */
    public function setMarginCoin($marginCoin): void
    {
        $this->marginCoin = $marginCoin;
    }

    /**
     * @return mixed
     */
    public function getSymbol()
    {
        return $this->symbol;
    }

    /**
     * @param mixed $symbol
     */
    public function setSymbol($symbol): void
    {
        $this->symbol = $symbol;
    }

    /**
     * @return mixed
     */
    public function getMarginMode()
    {
        return $this->marginMode;
    }

    /**
     * @param mixed $marginMode
     */
    public function setMarginMode($marginMode): void
    {
        $this->marginMode = $marginMode;
    }


}