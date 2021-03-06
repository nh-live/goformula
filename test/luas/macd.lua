--
-- Created by IntelliJ IDEA.
-- User: admin
-- Date: 2017/6/19
-- Time: 下午3:59
-- To change self template use File | Settings | File Templates.
--


MACDClass = {
}

MACDClass['name'] = "MACD"
MACDClass['argName'] = {'short', 'long', 'mid'}
MACDClass['short'] = {12, 2, 200}
MACDClass['long'] = {26, 2, 200}
MACDClass['mid'] = {9, 2, 200}
MACDClass['vars'] = {'DIF', 'DEA', 'MACD'}
MACDClass['noDraw'] = {0, 0, 0}
MACDClass['color'] = {{Red=255, Green=255, Blue=255}, {}, {}}
MACDClass['lineThick'] = {1, 1, 1}

function MACDClass:new(data, short, long, mid)
    o = {}
    setmetatable(o, self)
    self.__index = self

    o.data = data
    o.short = Scalar(short)
    o.long = Scalar(long)
    o.mid = Scalar(mid)
    o.close = CLOSE(data)
    o.const2 = Scalar(2)
    o.ema_close_short = EMA(o.close, o.short)
    o.ema_close_long = EMA(o.close, o.long)
    o.dif = SUB(o.ema_close_short, o.ema_close_long)
    o.dea = EMA(o.dif, o.mid)

    o.dif_sub_dea = SUB(o.dif, o.dea)
    o.macd = MUL(o.dif_sub_dea, o.const2)
    o.enter_long = CROSS(o.dif, o.dea)
    o.enter_short = CROSS(o.dea, o.dif)

    o.ref_values = {o.dif, o.dea, o.macd, o.enter_long, o.enter_short}

    return o
end

function MACDClass:updateLastValue()
    self.close.UpdateLastValue()
    self.ema_close_short.UpdateLastValue()
    self.ema_close_long.UpdateLastValue()
    self.dif.UpdateLastValue()
    self.dea.UpdateLastValue()
    self.dif_sub_dea.UpdateLastValue()
    self.macd.UpdateLastValue()
    self.enter_long.UpdateLastValue()
    self.enter_short.UpdateLastValue()
end

function MACDClass:Len()
    return self.data.Len()
end


function MACDClass:Get(index)
    return {
        self.dif.Get(index),
        self.dea.Get(index),
        self.macd.Get(index),
        self.enter_long.Get(index),
        self.enter_short.Get(index),
    }
end

FormulaClass = MACDClass
