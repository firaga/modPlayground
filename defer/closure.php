<?php
/**
 * Created by IntelliJ IDEA
 * User: jichen.zhou@eeoa.com
 * Date: 2021/6/22
 * Time: 10:51 上午
 */

class closure
{
    public function test()
    {
        $i = 1;
        $b = function () use ($i) {
            $c = $i;
            echo $c . PHP_EOL;
        };
        $b();
    }
}