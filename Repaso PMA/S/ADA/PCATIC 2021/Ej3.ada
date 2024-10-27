Procedure Ej3 is

--Vamos a tener task Empleado (uno solo)
-- Task embarazada
-- Task anciano
-- Task persona normal

Task type embarazada;

Task type persona_normal;

Task type anciano;

Task Empleado is

    Entry entryAnciano();
    Entry entryEmbarazada();
    Entry entryPersonaNormal();

End Empleado

arrPersonas: array (1..P) of persona_normal;
arrEmbarazadas: array (1..P) of embarazada;
arrAncianos: array (1..P) of anciano;

Task Body Empleado is

Begin
    loop
        SELECT
            when (entryEmbarazada'count = 0 and entryAnciano'count = 0) =>
                accept entryPersonaNormal() do
                    AtenderPedido();
                end entryPersonaNormal;
        OR
            when (entryEmbarazada'count > 0) =>
                accept entryEmbarazada() do
                    AtenderPedido();
                end entryEmbarazada;
        OR
            when (entryEmbarazada'count = 0 and entryAnciano'count > 0) =>
                accept entryAnciano() do
                    AtenderPedido();
                end entryAnciano;
        END SELECT;

    end loop;


End Empleado;


Task body persona_normal is

begin
    Empleado.entryPersonaNormal();
end persona_normal;

Task body embarazada is

begin
    SELECT
        Empleado.entryEmbarazada();
    ELSE
        null;
    END SELECT;

end embarazada;

Task body anciano is


begin
    SELECT
        Empleado.entryAnciano();
    OR DELAY 300.0;
        null;
    END SELECT;

end anciano;

begin

end Ej3;
