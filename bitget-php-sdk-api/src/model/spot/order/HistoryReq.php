<?php


namespace bitget\model\spot\order;


class HistoryReq
{
    /**
     * Currency pair
     */
    var string $symbol;
    /**
     * The orderId is passed in. The data before the orderId desc
     */
    var string $after;
    /**
     * Pass in the data after the orderId asc
     */
    var string $before;
    /**
     * Number of returned results Default 100, maximum 500
     */
    var string $limit;

    /**
     * @return string
     */
    public function getSymbol(): string
    {
        return $this->symbol;
    }

    /**
     * @param string $symbol
     */
    public function setSymbol(string $symbol): void
    {
        $this->symbol = $symbol;
    }

    /**
     * @return string
     */
    public function getAfter(): string
    {
        return $this->after;
    }

    /**
     * @param string $after
     */
    public function setAfter(string $after): void
    {
        $this->after = $after;
    }

    /**
     * @return string
     */
    public function getBefore(): string
    {
        return $this->before;
    }

    /**
     * @param string $before
     */
    public function setBefore(string $before): void
    {
        $this->before = $before;
    }

    /**
     * @return string
     */
    public function getLimit(): string
    {
        return $this->limit;
    }

    /**
     * @param string $limit
     */
    public function setLimit(string $limit): void
    {
        $this->limit = $limit;
    }


}